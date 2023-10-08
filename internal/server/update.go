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
		return nil, err
	}
	defer client.Close()

	if req.GetValue() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is blank")
	}
	v := req.GetValue()
	_, err = client.Put(ctx, "/key", v)
	if err != nil {
		return nil, err
	}

	return &pb.ServiceUpdateResponse{
		Result: "OK",
	}, nil
}

// エラーハンドリング
// 環境変数valult
