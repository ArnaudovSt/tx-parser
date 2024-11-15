package txparser

import (
	"github.com/ArnaudovSt/tx-parser/errors"
	"github.com/ArnaudovSt/tx-parser/storage"
	"github.com/ArnaudovSt/tx-parser/types"
)

type ITxParser interface {
	// last parsed block
	GetCurrentBlock() (uint64, error)
	// add address to observer
	Subscribe(address string) (bool, error)
	// remove address from observer
	Unsubscribe(address string) (bool, error)
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) ([]*types.Transaction, error)
}

var _ ITxParser = (*txParser)(nil)

type txParser struct {
	storage storage.IStorage
}

func NewTxParser(storage storage.IStorage) *txParser {
	return &txParser{
		storage: storage,
	}
}

func (p *txParser) GetCurrentBlock() (uint64, error) {
	b, err := p.storage.GetLatestBlock()
	if err != nil {
		return 0, err
	}

	if b == nil {
		return 0, errors.ErrServiceUnavailable
	}

	return b.Number.Uint64(), nil
}

func (p *txParser) GetTransactions(address string) ([]*types.Transaction, error) {
	txs, err := p.storage.GetTransactions(address)
	if err != nil {
		return nil, err
	}
	return txs, nil
}

func (p *txParser) Subscribe(address string) (bool, error) {
	err := p.storage.Subscribe(address)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *txParser) Unsubscribe(address string) (bool, error) {
	err := p.storage.Unsubscribe(address)
	if err != nil {
		return false, err
	}
	return true, nil
}
