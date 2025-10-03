package grpc

import (
	"fmt"
	"log"
	"net"

	"smart-dokan/usersvc/internal/ports"

	pb "github.com/asadlive84/smart-dokan-pb/golang/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api                               ports.APIPort
	port                              string
	pb.UnimplementedUserServiceServer // take this diffrence must
}

func NewAdapter(api ports.APIPort, port string) (*Adapter, error) {
	return &Adapter{
		api:  api,
		port: port,
	}, nil
}

func (a *Adapter) Run() {
	var err error

	fmt.Printf(" app port %+v\n", a.port)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", a.port))

	if err != nil {
		log.Fatalf("failed to listen on port %s, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryLoggerInterceptor),
	)

	pb.RegisterUserServiceServer(grpcServer, a)

	reflection.Register(grpcServer)

	log.Printf("Server run on the PORT: %s ", a.port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("faild to server grpc on port: ")
	}

}
