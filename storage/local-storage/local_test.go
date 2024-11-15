package localstorage

import (
	"math/big"
	"testing"

	"github.com/ArnaudovSt/tx-parser/errors"
	"github.com/ArnaudovSt/tx-parser/types"
	"github.com/stretchr/testify/assert"
)

func TestNewLocalStorage(t *testing.T) {
	reorgDepthLimit := uint64(10)
	storage := NewLocalStorage(reorgDepthLimit)

	assert.NotNil(t, storage)
	assert.Equal(t, reorgDepthLimit, storage.reorgDepthLimit)
}

func TestLocalStorage_GetLatestBlock(t *testing.T) {
	storage := NewLocalStorage(10)

	block := &types.Block{Number: big.NewInt(1)}
	storage.AppendBlock(block)

	block2 := &types.Block{Number: big.NewInt(2)}
	storage.AppendBlock(block2)

	latestBlock, err := storage.GetLatestBlock()
	assert.NoError(t, err)
	assert.Equal(t, block2, latestBlock)
}

func TestLocalStorage_GetTransactions(t *testing.T) {
	storage := NewLocalStorage(10)

	address := "0x123"
	storage.Subscribe(address)

	anotherAddress := "0x456"
	block := &types.Block{Number: big.NewInt(1), Hash: "blockHash1", Transactions: []*types.Transaction{{From: address}, {From: anotherAddress}}}
	storage.AppendBlock(block)

	txs, err := storage.GetTransactions(address)
	assert.NoError(t, err)
	assert.Len(t, txs, 1)
	assert.Equal(t, address, txs[0].From)
}

func TestLocalStorage_SubscribeAlreadyExists(t *testing.T) {
	storage := NewLocalStorage(10)
	address := "0x123"

	err := storage.Subscribe(address)
	assert.NoError(t, err)
	assert.Contains(t, storage.subscriptions, address)

	err = storage.Subscribe(address)
	assert.Error(t, err)
	assert.Equal(t, errors.ErrAlreadyExists, err)
}

func TestLocalStorage_Unsubscribe(t *testing.T) {
	storage := NewLocalStorage(10)
	address := "0x123"
	storage.Subscribe(address)

	err := storage.Unsubscribe(address)
	assert.NoError(t, err)
	assert.NotContains(t, storage.subscriptions, address)
}

func TestLocalStorage_AppendBlock(t *testing.T) {
	storage := NewLocalStorage(2)
	block1 := &types.Block{Number: big.NewInt(1)}
	block2 := &types.Block{Number: big.NewInt(2)}

	storage.AppendBlock(block1)
	storage.AppendBlock(block2)

	assert.Equal(t, block1, storage.chain.PeekTail())
	assert.Equal(t, block2, storage.chain.PeekHead())
}

func TestLocalStorage_AppendBlockDiscardsOldBlocks(t *testing.T) {
	storage := NewLocalStorage(2)

	sender1 := "0x1"
	block1 := &types.Block{Number: big.NewInt(1), Hash: "0x111", Transactions: []*types.Transaction{{From: sender1}}}
	storage.Subscribe(sender1)

	sender2 := "0x2"
	block2 := &types.Block{Number: big.NewInt(2), Hash: "0x222", Transactions: []*types.Transaction{{From: sender2}}}
	storage.Subscribe(sender2)

	sender3 := "0x3"
	block3 := &types.Block{Number: big.NewInt(3), Hash: "0x333", Transactions: []*types.Transaction{{From: sender3}}}
	storage.Subscribe(sender3)

	storage.AppendBlock(block1)
	storage.AppendBlock(block2)
	storage.AppendBlock(block3)

	txsSender1, err := storage.GetTransactions(sender1)
	assert.NoError(t, err)
	assert.Empty(t, txsSender1)

	txsSender2, err := storage.GetTransactions(sender2)
	assert.NoError(t, err)
	assert.Len(t, txsSender2, 1)
	assert.Equal(t, sender2, txsSender2[0].From)

	txsSender3, err := storage.GetTransactions(sender3)
	assert.NoError(t, err)
	assert.Len(t, txsSender3, 1)
	assert.Equal(t, sender3, txsSender3[0].From)

	assert.Equal(t, block2, storage.chain.PeekTail())
	assert.Equal(t, block3, storage.chain.PeekHead())
}

func TestLocalStorage_PopLatestBlock(t *testing.T) {
	storage := NewLocalStorage(10)

	sender1 := "0x1"
	block1 := &types.Block{Number: big.NewInt(1), Hash: "0x111", Transactions: []*types.Transaction{{From: sender1}}}
	storage.Subscribe(sender1)

	sender2 := "0x2"
	block2 := &types.Block{Number: big.NewInt(2), Hash: "0x222", Transactions: []*types.Transaction{{From: sender2}}}
	storage.Subscribe(sender2)

	storage.AppendBlock(block1)
	storage.AppendBlock(block2)

	poppedBlock, err := storage.PopLatestBlock()
	assert.NoError(t, err)
	assert.Equal(t, block2, poppedBlock)

	latestBlock, err := storage.GetLatestBlock()
	assert.NoError(t, err)
	assert.Equal(t, block1, latestBlock)

	txs, err := storage.GetTransactions(sender2)
	assert.NoError(t, err)
	assert.Empty(t, txs)
}
