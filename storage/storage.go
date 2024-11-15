package storage

import "github.com/ArnaudovSt/tx-parser/types"

// mockgen -source=storage.go -destination=mocks/mock_storage.go -package=mocks

type IReader interface {
	GetLatestBlock() (*types.Block, error)
	GetTransactions(address string) ([]*types.Transaction, error)
}

type IWriter interface {
	AtomicWrite(func(IAtomicWriter) error) error
	Subscribe(address string) error
	Unsubscribe(address string) error
}

type IAtomicWriter interface {
	AppendBlock(block *types.Block) error
	PopLatestBlock() (*types.Block, error)
}

type IStorage interface {
	IReader
	IWriter
}
