package server

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "bitbucket.org/zkrhm-fdn/microsvc-starter/kroto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Server struct server contains all methods needs for protobuf servering.
type Server struct {
}

//NewServer create server instance
func NewServer() *Server {
	return &Server{}
}

//SayHello the server side implementation
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello" + in.Name,
	}, nil
}

//Run running service on given port (on parameter)
func (s *Server) Run(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen : %v ", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	if err := s.Server(list); err != nil {
		log.Fatalf("Failed to serve : %s", err)
	}

	fmt.Println("> listening on port ", port)
}
