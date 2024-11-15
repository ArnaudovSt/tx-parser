package integration

import (
	"math/big"
	"testing"

	storage "github.com/ArnaudovSt/tx-parser/storage"
	localstorage "github.com/ArnaudovSt/tx-parser/storage/local-storage"
	txparser "github.com/ArnaudovSt/tx-parser/tx-parser"
	"github.com/ArnaudovSt/tx-parser/types"
	"github.com/stretchr/testify/assert"
)

func TestTxParser(t *testing.T) {
	s := localstorage.NewLocalStorage(2)
	parser := txparser.NewTxParser(s)

	sender1 := "0x1"
	block1 := &types.Block{Number: big.NewInt(1), Hash: "0x111", Transactions: []*types.Transaction{{From: sender1}}}
	success, err := parser.Subscribe(sender1)
	assert.NoError(t, err)
	assert.True(t, success)

	sender2 := "0x2"
	block2 := &types.Block{Number: big.NewInt(2), Hash: "0x222", Transactions: []*types.Transaction{{From: sender2}}}
	success, err = parser.Subscribe(sender2)
	assert.NoError(t, err)
	assert.True(t, success)

	sender3 := "0x3"
	block3 := &types.Block{Number: big.NewInt(3), Hash: "0x333", Transactions: []*types.Transaction{{From: sender3}}}
	success, err = parser.Subscribe(sender3)
	assert.NoError(t, err)
	assert.True(t, success)

	blocks := []*types.Block{block1, block2, block3}
	err = s.AtomicWrite(func(w storage.IAtomicWriter) error {
		for _, b := range blocks {
			err := w.AppendBlock(b)
			if err != nil {
				return err
			}
		}
		return nil
	})
	assert.NoError(t, err)

	blockNumber, err := parser.GetCurrentBlock()
	assert.NoError(t, err)
	assert.Equal(t, uint64(3), blockNumber)

	txsSender1, err := parser.GetTransactions(sender1)
	assert.NoError(t, err)
	assert.Empty(t, txsSender1)

	txsSender2, err := parser.GetTransactions(sender2)
	assert.NoError(t, err)
	assert.Len(t, txsSender2, 1)
	assert.Equal(t, sender2, txsSender2[0].From)

	txsSender3, err := parser.GetTransactions(sender3)
	assert.NoError(t, err)
	assert.Len(t, txsSender3, 1)
	assert.Equal(t, sender3, txsSender3[0].From)

	success, err = parser.Unsubscribe(sender2)
	assert.NoError(t, err)
	assert.True(t, success)

	success, err = parser.Unsubscribe(sender3)
	assert.NoError(t, err)
	assert.True(t, success)

	txsSender2, err = parser.GetTransactions(sender2)
	assert.NoError(t, err)
	assert.Empty(t, txsSender2)

	txsSender3, err = parser.GetTransactions(sender3)
	assert.NoError(t, err)
	assert.Empty(t, txsSender3)
}
