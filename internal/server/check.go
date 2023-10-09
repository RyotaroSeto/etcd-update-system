package server

import (
	"context"
	"etcd-update-system/internal/etcd"
	pb "etcd-update-system/pkg/gen/something/v1"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *systemServer) Check(ctx context.Context, req *pb.ServiceCheckRequest) (*pb.ServiceCheckResponse, error) {
	client, err := etcd.NewClient()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create etcd client: %v", err)
	}
	defer func() {
		_ = client.Close()
	}()

	res, err := client.Get(ctx, "/key")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update value in etcd: %v", err)
	}
	if res.Count == 0 {
		return nil, status.Errorf(codes.NotFound, "/key not found")
	}
	log.Println(string(res.Kvs[0].Value))

	return &pb.ServiceCheckResponse{
		Result: string(res.Kvs[0].Value),
	}, nil
}
