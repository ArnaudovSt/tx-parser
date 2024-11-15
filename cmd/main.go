package main

import (
	"context"
	"log"

	"github.com/ArnaudovSt/tx-parser/api"
	evmclient "github.com/ArnaudovSt/tx-parser/client/evm-client"
	"github.com/ArnaudovSt/tx-parser/config"
	pollingprovider "github.com/ArnaudovSt/tx-parser/provider/polling-provider"
	localstorage "github.com/ArnaudovSt/tx-parser/storage/local-storage"
	txparser "github.com/ArnaudovSt/tx-parser/tx-parser"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("[ERROR] Failed to load .env file")
		return
	}

	config, err := config.GetConfig()
	if err != nil {
		return
	}

	storage := localstorage.NewLocalStorage(config.ReorgDepthLimit)
	client := evmclient.NewEVMClient(config.EvmRpcUrl)
	provider := pollingprovider.NewPollingProvider(client, storage, config.AvgBlockTime)
	parser := txparser.NewTxParser(storage)
	server := api.NewServer(parser, config.ServerADDR)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := provider.Start(ctx); err != nil {
			log.Println("[ERROR] Failed to start provider:", err)
		}
	}()

	if err := server.Start(); err != nil {
		log.Println("[ERROR] Failed to start server:", err)
		cancel()
	}
}
