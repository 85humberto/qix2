package main

import (
	"context"
	"encoding/json"
	"flag"
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
	// addressSerQix = "server1:50051"
	portSerLoad = ":50052"
)

var wg sync.WaitGroup

type serverLoad struct {
	pb.UnimplementedLoadServer
}

func (s *serverLoad) EnviaLoad(ctx context.Context, in *pb.LoadInfo) (*pb.LoadRecebido, error) {
	ModificaLoad(in.Servidor, in.Load)
	log.Printf("Servidor %v com %d", in.Servidor, in.Load)
	return &pb.LoadRecebido{Sucesso: true}, nil
}

type Qix struct {
	Transação    string `json:"transacao"`
	Complexidade int    `json:"complexidade"`
}

type Backend struct {
	URL    string
	Load   int64
	Status bool
}

type ServerPool struct {
	backends []*Backend
}

func (s *ServerPool) AddBackend(backend *Backend) {
	s.backends = append(s.backends, backend)
}

func ModificaLoad(server string, load int64) {
	for _, s := range serverpool.backends {
		if s.URL == server {
			s.Load = load
		}
	}
}

func (b *Backend) ModificaStatus(s bool) {
	b.Status = s
}

//Pega o servidor com o menor Load e retorna a URL
func ProximoServer() string {
	v := serverpool.backends[0].Load
	p := serverpool.backends[0].URL
	// for _, s := range serverpool.backends {
	// 	if s.Load <= int64(v) {
	// 		p = s.URL
	// 	}
	// }
	for i := 1; i < len(serverpool.backends); i++ {
		if serverpool.backends[i].Load <= v {
			p = serverpool.backends[i].URL
			// } else {
			// 	p = serverpool.backends[i+1].URL
		}
	}
	return p
}

var serverpool ServerPool

func main() {
	// define backends
	flag.Parse()
	servers := flag.Args()
	for _, s := range servers {
		b := Backend{
			URL:    s,
			Status: true,
		}
		serverpool.AddBackend(&b)
	}

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

	ps := ProximoServer() + ":50051"
	fmt.Println(ps)
	//grpc
	conn, err := grpc.Dial(ps, grpc.WithInsecure(), grpc.WithBlock())
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
