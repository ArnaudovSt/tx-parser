package types

type BlockResponse struct {
	JSONRPC string                 `json:"jsonrpc"`
	ID      int                    `json:"id"`
	Result  *Block                 `json:"result"`
	Error   map[string]interface{} `json:"error"`
}
