package server

import (
	pb "etcd-update-system/pkg/gen/something/v1"
)

type systemServer struct {
	pb.UnimplementedServiceServer
}

var _ pb.ServiceServer = &systemServer{}

func NewSystemServer() *systemServer {
	return &systemServer{}
}
