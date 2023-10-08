package main

import (
	"fmt"
	"log"
	"net"
	"os"
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

func run() error {
	port := 8585
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterServiceServer(s, server.NewSystemServer())

	go func() {
		log.Printf("start gRPC server port: %v", port)
		s.Serve(l)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()

	log.Println("clean shutdown")
	return nil
}
