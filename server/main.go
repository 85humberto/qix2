package main

import (
	"context"
	"flag"
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

func (s *serverQix) EnviaQix(ctx context.Context, in *pb.QixRequest) (*pb.QixReply, error) {
	if (load + int(in.Complexidade)) > 100 {
		log.Println("Transação rejeitada, servidor sobrecarregado")
		return &pb.QixReply{Sucesso: false}, nil
	} else {
		load += int(in.Complexidade)
		log.Printf("Transação recebida: %v %d", in.Transacao, in.Complexidade)
		return &pb.QixReply{Servidor: hostname, Sucesso: true}, nil
	}
}

func trabalha() {
	defer wg.Done()
	for {
		log.Printf("Load atual %d\n", load)
		time.Sleep(1 * time.Second)
		if load > 0 {
			load -= 1
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
	//inicia das gorouteines
	wg.Add(1)
	go trabalha()
	wg.Add(1)
	go enviaLoad()
	//abre porta pra escutar
	lis, err := net.Listen("tcp", portSerQix)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//cria servidor grpc
	s := grpc.NewServer()
	pb.RegisterQixServer(s, &serverQix{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
