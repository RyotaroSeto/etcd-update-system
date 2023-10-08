package server

import (
	"context"
	"etcd-update-system/internal/etcd"
	pb "etcd-update-system/pkg/gen/something/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *systemServer) Update(ctx context.Context, req *pb.ServiceUpdateRequest) (*pb.ServiceUpdateResponse, error) {
	client, err := etcd.NewClient()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create etcd client: %v", err)
	}
	defer func() {
		_ = client.Close()
	}()

	v := req.GetValue()
	if v == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is blank")
	}

	_, err = client.Put(ctx, "/key", v)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update value in etcd: %v", err)
	}

	return &pb.ServiceUpdateResponse{
		Result: "OK",
	}, nil
}
