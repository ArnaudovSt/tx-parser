package types

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Block struct {
	Number       *big.Int       `json:"number"`
	Hash         string         `json:"hash"`
	ParentHash   string         `json:"parentHash"`
	Transactions []*Transaction `json:"transactions"`
	// TODO: add all block fields
}

// Custom unmarshalling for Block to handle hex string to *big.Int for Number
func (b *Block) UnmarshalJSON(data []byte) error {
	// Define a separate struct with Number as string to parse JSON initially
	type Alias Block
	aux := &struct {
		Number string `json:"number"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	// Unmarshal into the temporary structure
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Convert the number from hex string to *big.Int
	num := new(big.Int)
	num, ok := num.SetString(aux.Number[2:], 16) // Skip "0x" prefix for hex string
	if !ok {
		return fmt.Errorf("invalid block number format")
	}
	b.Number = num

	return nil
}
