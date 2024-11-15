package evmclient

import (
	"context"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ArnaudovSt/tx-parser/types"
	"github.com/stretchr/testify/assert"
)

func TestNewEVMClient(t *testing.T) {
	url := "http://localhost:8545"
	client := NewEVMClient(url)

	assert.NotNil(t, client)
	assert.Equal(t, url, client.url)
}

func TestGetBlockByHash(t *testing.T) {
	expectedBlock := &types.Block{Number: big.NewInt(1), Hash: "0x2", ParentHash: "0x3", Transactions: []*types.Transaction{{From: "0x4", To: "0x5", TransactionHash: "0x6"}}}
	expectedRequestBody := `{"jsonrpc":"2.0","method":"eth_getBlockByHash","params":["0x123", true],"id":1}`

	var actualRequestBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		actualRequestBody = string(bodyBytes)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"jsonrpc":"2.0","result":{"number":"0x1","hash":"0x2","parentHash":"0x3","transactions":[{"from":"0x4","to":"0x5","hash":"0x6"}]},"id":1}`))
	}))
	defer server.Close()

	client := NewEVMClient(server.URL)
	block, err := client.GetBlockByHash(context.Background(), "0x123")

	assert.NoError(t, err)
	assert.Equal(t, expectedBlock, block)
	assert.JSONEq(t, expectedRequestBody, actualRequestBody)

}

func TestGetLatestBlock(t *testing.T) {
	expectedBlock := &types.Block{Number: big.NewInt(1)}
	expectedRequestBody := `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest", true],"id":1}`

	var actualRequestBody string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		actualRequestBody = string(bodyBytes)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"jsonrpc":"2.0","result":{"number":"0x1"},"id":1}`))
	}))
	defer server.Close()

	client := NewEVMClient(server.URL)
	block, err := client.GetLatestBlock(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, expectedBlock, block)
	assert.JSONEq(t, expectedRequestBody, actualRequestBody)
}
