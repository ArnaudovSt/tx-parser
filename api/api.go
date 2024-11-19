package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ArnaudovSt/tx-parser/errors"
	txparser "github.com/ArnaudovSt/tx-parser/tx-parser"
	"github.com/ArnaudovSt/tx-parser/types"
)

type server struct {
	txParser txparser.ITxParser
	addr     string
}

func NewServer(txParser txparser.ITxParser, addr string) *server {
	return &server{txParser: txParser, addr: addr}
}

func (s *server) Start() error {
	http.HandleFunc("/blocks", s.BlocksHandler)
	http.HandleFunc("/subscriptions", s.SubscriptionsHandler)
	http.HandleFunc("/transactions", s.TransactionsHandler)

	fmt.Printf("Server started on %s\n", s.addr)
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return err
	}
	return nil
}

func (s *server) BlocksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		block, err := s.txParser.GetCurrentBlock()
		if err != nil {
			if errors.IsUnavailable(err) {
				log.Printf("[ERROR] Current block is not initialised")
				http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
				return
			}

			log.Printf("[ERROR] Failed to get current block: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		log.Printf("[DEBUG] Request handled successfully for latest block: %d", block)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]uint64{"result": block})

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (s *server) SubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var requestData struct {
			Address string `json:"address"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		success, err := s.txParser.Subscribe(requestData.Address)
		if err != nil || !success {
			if errors.IsAlreadyExists(err) {
				http.Error(w, "Address is already subscribed", http.StatusConflict)
				return
			}

			log.Printf("[ERROR] Failed to subscribe: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		log.Printf("[DEBUG] Subscription request handled successfully for address: %s", requestData.Address)

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"result": "Subscription created successfully",
		})

	case http.MethodDelete:
		var requestData struct {
			Address string `json:"address"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		success, err := s.txParser.Unsubscribe(requestData.Address)
		if err != nil || !success {
			log.Printf("[ERROR] Failed to unsubscribe: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		log.Printf("[DEBUG] Unsubscribe request handled successfully for address: %s", requestData.Address)

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (s *server) TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		address := r.URL.Query().Get("address")

		transactions, err := s.txParser.GetTransactions(address)
		if err != nil {
			log.Printf("[ERROR] Failed to get transactions: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if len(transactions) == 0 {
			log.Printf("[DEBUG] No transactions found for address: %s", address)
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]string{"result": "[]"})
			return

		}

		log.Printf("[DEBUG] Get transactions request handled successfully for address: %s", address)

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string][]*types.Transaction{"result": transactions})

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
