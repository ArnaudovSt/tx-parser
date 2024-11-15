package localstorage

import (
	"strings"
	"sync"

	"github.com/ArnaudovSt/tx-parser/errors"
	"github.com/ArnaudovSt/tx-parser/storage"
	"github.com/ArnaudovSt/tx-parser/types"
)

var _ storage.IStorage = (*localStorage)(nil)

type blockTransactions map[string][]*types.Transaction

// address -> blockHash -> transactions
type subscriptions map[string]blockTransactions

type localStorage struct {
	chain           *types.Chain
	subscriptions   subscriptions
	reorgDepthLimit uint64
	mu              sync.RWMutex
}

func NewLocalStorage(reorgDepthLimit uint64) *localStorage {
	return &localStorage{
		chain:           types.NewChain(),
		subscriptions:   make(subscriptions),
		reorgDepthLimit: reorgDepthLimit,
	}
}

func (l *localStorage) GetLatestBlock() (*types.Block, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.chain.PeekHead(), nil
}

func (l *localStorage) GetTransactions(address string) ([]*types.Transaction, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var txs []*types.Transaction
	for _, blockTxs := range l.subscriptions[strings.ToLower(address)] {
		txs = append(txs, blockTxs...)
	}

	return txs, nil
}

func (l *localStorage) Subscribe(address string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	address = strings.ToLower(address)

	if _, ok := l.subscriptions[address]; ok {
		return errors.ErrAlreadyExists
	}

	l.subscriptions[address] = make(map[string][]*types.Transaction)
	return nil
}

func (l *localStorage) Unsubscribe(address string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.subscriptions, strings.ToLower(address))
	return nil
}

// Helper method to perform atomic write operations on the storage.
func (l *localStorage) AtomicWrite(operation func(w storage.IAtomicWriter) error) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	return operation(l)
}

func (l *localStorage) AppendBlock(block *types.Block) error {
	l.chain.Append(block)
	l.addTransactions(block)

	earliestBlock := l.chain.PeekTail()
	if block.Number.Uint64()-earliestBlock.Number.Uint64() >= l.reorgDepthLimit {
		_ = l.chain.PopTail()
		l.removeTransactions(earliestBlock)
	}

	return nil
}

func (l *localStorage) PopLatestBlock() (*types.Block, error) {
	block := l.chain.PopHead()
	l.removeTransactions(block)
	return block, nil
}

func (l *localStorage) addTransactions(block *types.Block) {
	for _, tx := range block.Transactions {
		blockHash := strings.ToLower(block.Hash)
		txFrom := strings.ToLower(tx.From)
		if _, ok := l.subscriptions[txFrom]; ok {
			l.subscriptions[txFrom][blockHash] = append(l.subscriptions[txFrom][blockHash], tx)
		}

		txTo := strings.ToLower(tx.To)
		if _, ok := l.subscriptions[txTo]; ok {
			l.subscriptions[txTo][blockHash] = append(l.subscriptions[txTo][blockHash], tx)
		}
	}
}

func (l *localStorage) removeTransactions(block *types.Block) {
	for address := range l.subscriptions {
		delete(l.subscriptions[strings.ToLower(address)], strings.ToLower(block.Hash))
	}
}
