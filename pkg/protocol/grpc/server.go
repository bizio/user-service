package grpc

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/bizio/user-service/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, v1Api v1.UserServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	v1.RegisterUserServiceServer(server, v1Api)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			// sig is a ^C, handle it
			log.Println("Shutting down User gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Printf("Starting User gRPC server on port %s...\n", port)
	return server.Serve(listen)

}
