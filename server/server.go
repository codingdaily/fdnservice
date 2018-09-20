package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"go.uber.org/zap/zapcore"

	pb "bitbucket.org/zkrhm-fdn/microsvc-starter/kroto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	logger *zap.Logger
)

func init() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ = config.Build()
}

//Server struct server contains all methods needs for protobuf servering.
type Server struct {
}

//NewServer create server instance
func NewServer() *Server {
	return &Server{}
}

//SayHello the server side implementation
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	sugar := logger.Sugar()
	sugar.Infow("Accepting request ",
		"param", in.Name,
	)
	return &pb.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

//Run running service on given port (on parameter)
func (s *Server) Run(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to Listen : %v ", err)
	}

	grpcSvr := grpc.NewServer()
	pb.RegisterGreeterServer(grpcSvr, NewServer())
	reflection.Register(grpcSvr)

	fmt.Println("> listening on port ", port)
	if err := grpcSvr.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %s", err)
	}

}
