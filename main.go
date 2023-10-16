package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os/signal"
	"syscall"

	pb "etcd-update-system/pkg/gen/something/v1"

	"etcd-update-system/internal/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}
}

const port = 8585

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	srv := grpc.NewServer()
	reflection.Register(srv)
	pb.RegisterServiceServer(srv, server.NewSystemServer())

	go func() {
		<-ctx.Done()
		srv.GracefulStop()
	}()

	err = srv.Serve(l)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}

	return nil
}
