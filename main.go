package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"
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
	server := grpc.NewServer()
	// pb.RegisterGreeterServer(grpc, &gSrv)
	// client, err := etcd.NewClient()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// server, err := newServer(serverConfig{
	// 	Address: address,
	// 	DB:      pool,
	// 	etcd:    esClient,
	// })

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		s := <-quit
		log.Printf("got signal %v, attempting graceful shutdown", s)
		server.GracefulStop()
		wg.Done()
	}()

	log.Printf("start gRPC server port: %v", port)
	err = server.Serve(l)
	if err != nil {
		log.Fatalf("could not serve: %v", err)
	}

	wg.Wait()
	log.Println("clean shutdown")
	return nil
}
