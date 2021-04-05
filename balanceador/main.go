package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	pb "github.com/85humberto/qix2/pb"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
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

	//grpc
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewQixClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tq, err := c.EnviaQix(ctx, &pb.QixRequest{Transacao: q.Transação, Complexidade: int32(q.Complexidade)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Envio %t", tq.GetSucesso())
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(q)
}

// curl -H 'Content-Type: application/json' -d '{"transacao":"pagamento","complexidade":9}' -X POST http://localhost:8080/qix
