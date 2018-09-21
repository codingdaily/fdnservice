package server

import (
	"context"
	"fmt"
	"net"

	raven "github.com/getsentry/raven-go"

	pb "bitbucket.org/zkrhm-fdn/microsvc-starter/kroto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	appName = "fdnsvc"
)

var (
	logger *zap.Logger
)

// func init() {
// 	logger, _ = logconfig.NewZapLogger(appName)
// }

//Server struct server contains all methods needs for protobuf servering.
type Server struct {
	appName string
	logger  *zap.Logger
}

//NewServerWithLogger create server instance with zap logger parameter passed.
func NewServerWithLogger(loggerParam *zap.Logger) *Server {
	return &Server{appName: appName, logger: loggerParam}
}

func newServer() *Server {
	return &Server{}
}

//SayHello the server side implementation
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello " + in.Name,
	}, nil
}

//Run running service on given port (on parameter)
func (s *Server) Run(port string) {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		s.logger.Fatal(fmt.Sprint("Failed to Listen : %v ", err))
	}

	grpcSvr := grpc.NewServer()
	pb.RegisterGreeterServer(grpcSvr, newServer())
	reflection.Register(grpcSvr)

	s.logger.Info(fmt.Sprint("> listening on port ", port))
	fmt.Println("fmt> listening on port ", port)
	if err := grpcSvr.Serve(lis); err != nil {
		raven.CaptureErrorAndWait(err, nil)
		s.logger.Fatal(fmt.Sprint("Failed to serve : %s", err))
	}

}
