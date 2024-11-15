// The redisStorage is designed as a future improvement to utilize Redis as a storage, instead of the current in-memory storage.
// This could bring various benefits, such as persistence, better scalability, different eviction policies, etc. depending on the needs of the application.
package redisstorage

import (
	"github.com/ArnaudovSt/tx-parser/storage"
	"github.com/ArnaudovSt/tx-parser/types"
)

var _ storage.IStorage = (*redisStorage)(nil)

type redisStorage struct{}

func NewRedisStorage() *redisStorage {
	return &redisStorage{}
}

func (r *redisStorage) AppendBlock(block *types.Block) error {
	panic("unimplemented")
}

func (r *redisStorage) AtomicWrite(func(storage.IAtomicWriter) error) error {
	panic("unimplemented")
}

func (r *redisStorage) GetLatestBlock() (*types.Block, error) {
	panic("unimplemented")
}

func (r *redisStorage) GetTransactions(address string) ([]*types.Transaction, error) {
	panic("unimplemented")
}

func (r *redisStorage) PopLatestBlock() (*types.Block, error) {
	panic("unimplemented")
}

func (r *redisStorage) Subscribe(address string) error {
	panic("unimplemented")
}

func (r *redisStorage) Unsubscribe(address string) error {
	panic("unimplemented")
}
