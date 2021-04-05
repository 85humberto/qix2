package main

import (
	"context"
	"log"
	"net"

	pb "github.com/85humberto/qix2/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedQixServer
}

func (s *server) EnviaQix(ctx context.Context, in *pb.QixRequest) (*pb.QixReply, error) {
	log.Printf("Transação recebida: %v %d", in.Transacao, in.Complexidade)
	return &pb.QixReply{Sucesso: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterQixServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
