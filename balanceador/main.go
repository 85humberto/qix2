package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"

	pb "github.com/85humberto/qix2/pb"
	"google.golang.org/grpc"
)

const (
	addressSerQix = "localhost:50051"
	portSerLoad   = ":50052"
)

var wg sync.WaitGroup

// type serverQix struct {
// 	pb.UnimplementedQixServer
// }

type serverLoad struct {
	pb.UnimplementedLoadServer
}

func (s *serverLoad) EnviaLoad(ctx context.Context, in *pb.LoadInfo) (*pb.LoadRecebido, error) {
	log.Printf("Servidor %v com %d", in.Servidor, in.Load)
	return &pb.LoadRecebido{Sucesso: true}, nil
}

type Qix struct {
	Transação    string `json:"transacao"`
	Complexidade int    `json:"complexidade"`
}

func main() {
	defer wg.Done()
	wg.Add(1)
	go func() {
		lis, err := net.Listen("tcp", portSerLoad)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterLoadServer(s, &serverLoad{})
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	router := mux.NewRouter()
	router.HandleFunc("/qix", createQix).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func createQix(w http.ResponseWriter, r *http.Request) {
	var q Qix
	json.NewDecoder(r.Body).Decode(&q)
	fmt.Println("Transação recebida: ", q)

	//grpc
	conn, err := grpc.Dial(addressSerQix, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Não foi possível conectar: %v", err)
	}
	defer conn.Close()
	c := pb.NewQixClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tq, err := c.EnviaQix(ctx, &pb.QixRequest{Transacao: q.Transação, Complexidade: int32(q.Complexidade)})
	if err != nil {
		log.Fatalf("Não foi possível transmitir para o server: %v", err)
	}
	log.Printf("Envio %t", tq.GetSucesso())
}

// curl -H 'Content-Type: application/json' -d '{"transacao":"pagamento","complexidade":9}' -X POST http://localhost:8080/qix
