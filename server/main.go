package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "github.com/85humberto/qix2/pb"
	"google.golang.org/grpc"
)

const (
	portSerQix     = ":50051"
	addressSerLoad = "balanceador:50052"
)

var hostname string
var load = 0
var wg sync.WaitGroup

type serverQix struct {
	pb.UnimplementedQixServer
}

// type serverLoad struct {
// 	pb.UnimplementedLoadServer
// }

func (s *serverQix) EnviaQix(ctx context.Context, in *pb.QixRequest) (*pb.QixReply, error) {
	load += int(in.Complexidade)
	log.Printf("Transação recebida: %v %d", in.Transacao, in.Complexidade)
	return &pb.QixReply{Sucesso: true}, nil
}

func trabalha() {
	defer wg.Done()
	for {
		fmt.Printf("Load atual %d\n", load)
		time.Sleep(1 * time.Second)
		if load > 0 {
			load--
		}
	}
}

func enviaLoad() {
	defer wg.Done()
	for {
		time.Sleep(10 * time.Second)
		conn, err := grpc.Dial(addressSerLoad, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Não foi possível conectar: %v", err)
		}
		defer conn.Close()
		c := pb.NewLoadClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		tl, err := c.EnviaLoad(ctx, &pb.LoadInfo{Servidor: hostname, Load: int64(load)})
		if err != nil {
			log.Fatalf("Não foi possível transmitir para o server: %v", err)
		}
		log.Printf("Envio %t", tl.GetSucesso())
	}
}

func main() {
	// define hostname
	flag.Parse()
	hostname = flag.Arg(0)

	//gerenciamento das gorouteines
	wg.Add(1)
	go trabalha()
	wg.Add(1)
	go enviaLoad()

	lis, err := net.Listen("tcp", portSerQix)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQixServer(s, &serverQix{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
