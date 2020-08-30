package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// RunServer :
func RunServer() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	port := os.Getenv("PORT")

	router := mux.NewRouter()
	router.HandleFunc("/chain", GetBlockchainHandler).Methods("GET")
	router.HandleFunc("/add", AddBlockHandler).Methods("POST")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return err
	}

	return nil
}

// GetBlockchainHandler :
func GetBlockchainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bc)

}

// AddBlockHandler :
func AddBlockHandler(w http.ResponseWriter, r *http.Request) {

	t := Transaction{}

	json.NewDecoder(r.Body).Decode(&t)
	if t.IsTransactionSet() {
		w.Header().Set("Content-Type", "application/json")
		bc.AddBlock(t)
		json.NewEncoder(w).Encode(Message{
			Message: "the block was added successfuly :)",
		})
	} else {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Message{
			Message: "something went wrong  :( ... ",
		})
	}

}
