package client

import (
	"context"

	"github.com/ArnaudovSt/tx-parser/types"
)

// mockgen -source=client.go -destination=mocks/mock_client.go -package=mocks

type IClient interface {
	GetLatestBlock(ctx context.Context) (*types.Block, error)
	GetBlockByHash(ctx context.Context, hash string) (*types.Block, error)
}
