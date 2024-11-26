package evmclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ArnaudovSt/tx-parser/client"
	"github.com/ArnaudovSt/tx-parser/types"
)

var _ client.IClient = (*evmClient)(nil)

type evmClient struct {
	httpClient *http.Client
	url        string
}

func NewEVMClient(url string) *evmClient {
	return &evmClient{
		httpClient: &http.Client{},
		url:        url,
	}
}

func (e *evmClient) GetBlockByHash(ctx context.Context, hash string) (*types.Block, error) {
	jsonRequest := fmt.Sprintf(`{"jsonrpc":"2.0","method":"eth_getBlockByHash","params":["%s",true],"id":1}`, hash)
	return e.makeRequest(ctx, jsonRequest)
}

func (e *evmClient) GetLatestBlock(ctx context.Context) (*types.Block, error) {
	jsonRequest := `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["latest",true],"id":1}`
	return e.makeRequest(ctx, jsonRequest)
}

func (e *evmClient) makeRequest(ctx context.Context, jsonRequest string) (*types.Block, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "POST", e.url, strings.NewReader(jsonRequest))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := e.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var rpcResponse types.BlockResponse
	err = json.Unmarshal(body, &rpcResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	if rpcResponse.Error != nil {
		return nil, fmt.Errorf("JSON-RPC Error: %v", rpcResponse.Error)
	}

	return rpcResponse.Result, nil
}
