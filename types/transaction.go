package types

type Transaction struct {
	From            string `json:"from"`
	To              string `json:"to"`
	TransactionHash string `json:"hash"`
	// TODO: add all transaction fields
}
