package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gorilla/mux"

	pb "github.com/85humberto/qix2/pb"
	"google.golang.org/grpc"
)

const (
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
	URL     string
	Load    int64
	Status  bool
	Counter int
}

//funções de ordenação
type ByLoad []*Backend

func (a ByLoad) Len() int           { return len(a) }
func (a ByLoad) Less(i, j int) bool { return a[i].Load < a[j].Load }
func (a ByLoad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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
			//zera contador quando recebe um load.
			s.Counter = 0
		}
	}
}

func (b *Backend) ModificaStatus(s bool) {
	b.Status = s
}

func ProximoServer() string {
	sort.Sort(ByLoad(serverpool.backends))
	VerificaStatus()
	for i := 0; i < len(serverpool.backends); i++ {
		if serverpool.backends[i].Status {
			return serverpool.backends[i].URL
		}
	}
	return ""
}

func VerificaStatus() {
	for i := 0; i < len(serverpool.backends); i++ {
		if serverpool.backends[i].Counter >= 3 {
			serverpool.backends[i].Status = false
			log.Printf("Sem comunicação com o servidor %s", serverpool.backends[i].URL)
		} else {
			serverpool.backends[i].Status = true
		}
	}
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
	//goroutina do servidor grpc para receber informações de load dos servers
	go func() {
		//abre porta pra escutar
		lis, err := net.Listen("tcp", portSerLoad)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		//cria servidor grpc
		s := grpc.NewServer()
		pb.RegisterLoadServer(s, &serverLoad{})
		if err := s.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
		}
	}()

	wg.Add(1)
	//goroutina para incrementar o contador dos servers a cada 10 segundos.
	go func() {
		for {
			time.Sleep(10 * time.Second)
			for i := 0; i < len(serverpool.backends); i++ {
				serverpool.backends[i].Counter += 1
			}
		}
	}()
	//criar servidor http e define a rota /qix
	router := mux.NewRouter()
	router.HandleFunc("/qix", createQix).Methods("POST")
	log.Print(http.ListenAndServe(":8080", router))
}

func createQix(w http.ResponseWriter, r *http.Request) {
	var q Qix
	json.NewDecoder(r.Body).Decode(&q)
	log.Println("Transação recebida: ", q)
	//define próximo servidor
	ps := ProximoServer()
	if ps != "" {
		ps += ":50051"
	} else {
		log.Println("Nenhum servidor ativo no momento.")
	}
	//Enviar a transação e o load para o server com menor carga
	conn, err := grpc.Dial(ps, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Duration(5)*time.Second))
	if err != nil {
		log.Printf("Não foi possível conectar: %v", err)
		return
	}
	defer conn.Close()
	c := pb.NewQixClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tq, err := c.EnviaQix(ctx, &pb.QixRequest{Transacao: q.Transação, Complexidade: int32(q.Complexidade)})
	if err != nil {
		log.Printf("Não foi possível transmitir para o server: %v", err)
		return
	}
	log.Printf("Recebido pelo servidor %v: %t", tq.Servidor, tq.GetSucesso())
}
