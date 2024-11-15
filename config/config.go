package config

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	EvmRpcUrl       string
	AvgBlockTime    time.Duration
	ReorgDepthLimit uint64
	ServerADDR      string
}

func GetConfig() (*Config, error) {
	url := os.Getenv("EVM_RPC_URL")
	if url == "" {
		errMsg := "[ERROR] EVM_RPC_URL environment variable is required"
		log.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	avgBlockTime, err := time.ParseDuration(os.Getenv("AVG_BLOCK_TIME"))
	if err != nil {
		log.Printf("[ERROR] Failed to parse AVG_BLOCK_TIME environment variable: %v", err)
		return nil, err
	}

	reorgDepthLimit, err := strconv.ParseUint(os.Getenv("REORG_DEPTH_LIMIT"), 10, 64)
	if err != nil {
		log.Printf("[ERROR] Failed to parse REORG_DEPTH_LIMIT environment variable: %v", err)
		return nil, err
	}

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = ":8080"
	}

	return &Config{
		EvmRpcUrl:       url,
		AvgBlockTime:    avgBlockTime,
		ReorgDepthLimit: reorgDepthLimit,
		ServerADDR:      serverAddr,
	}, nil
}
