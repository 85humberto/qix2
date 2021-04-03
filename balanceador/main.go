package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Qix struct {
	Transação    string `json:"transacao"`
	Complexidade int    `json:"complexidade"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/qix", createQix).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func createQix(w http.ResponseWriter, r *http.Request) {
	var q Qix
	json.NewDecoder(r.Body).Decode(&q)

	fmt.Println("Transação recebida: ", q)

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(q)
}

// curl -H 'Content-Type: application/json' -d '{"transacao":"pagamento","complexidade":9}' -X POST http://localhost:8080/qix
