package server

import (
	"context"
	"etcd-update-system/internal/etcd"
	pb "etcd-update-system/pkg/gen/something/v1"
)

func (s *systemServer) Update(ctx context.Context, req *pb.ServiceUpdateRequest) (*pb.ServiceUpdateResponse, error) {
	client, err := etcd.NewClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	_, err = client.Put(ctx, "/key", "yes")
	if err != nil {
		return nil, err
	}
	return &pb.ServiceUpdateResponse{
		Message: "OK",
	}, nil
}

// errだとかえせないからどうにかする
// リクエスト受け取る
// 環境変数valult
