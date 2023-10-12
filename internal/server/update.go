package server

import (
	"context"
	"etcd-update-system/internal/etcd"
	pb "etcd-update-system/pkg/gen/something/v1"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
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
		// 以下はエラーレスポンス方法2
		st := status.New(codes.InvalidArgument, "Name is blank")
		v := &errdetails.BadRequest{
			FieldViolations: []*errdetails.BadRequest_FieldViolation{
				{
					Field:       "username",
					Description: "should not empty",
				},
			},
		}
		dt, _ := st.WithDetails(v)
		return nil, dt.Err()
		// 以下はclientでの取り出し方
		// st := status.Convert(err)
		// for _, detail := range st.Details() {
		// 	switch t := detail.(type) {
		// 	case *errdetails.BadRequest:
		// 		fmt.Println("Oops! Your request was rejected by the server.")
		// 		for _, violation := range t.GetFieldViolations() {
		// 			fmt.Printf("The %q field was wrong:\n", violation.GetField())
		// 			fmt.Printf("\t%s\n", violation.GetDescription())
		// 		}
		// 	}
		// }
		// 以下はエラーレスポンス方法1
		// return nil, status.Errorf(codes.InvalidArgument, "Name is blank")
		// 以下はエラーレスポンス方法3 protoに独自のエラーメッセージを用意している場合
		// st := status.New(codes.InvalidArgument, "some error occurred")
		// dt, _ := st.WithDetails(&pb.ErrorDetail{Code: pb.ErrorCode_EXPIRED_RECEIPT})
		// return nil, dt.Err()
		// 以下はclientでの取り出し方
		// st, _ := status.FromError(err)
		// for _, detail := range st.Details() {
		// 	switch t := detail.(type) {
		// 	case *errdetails.BadRequest:
		// 		fmt.Println("handle BadRequest case")
		// 	case *errdetails.QuotaFailure:
		// 		fmt.Println("handle QuotaFailure case")
		// 	case *pb.ErrorDetail:
		// 		// handle original error code
		// 		fmt.Println("error code:", t.Code)
		// 	}
		// }
		// 以下サンプル
		// https://github.com/jun06t/grpc-sample/blob/master/error-details/server/main.go
	}

	_, err = client.Put(ctx, "/key", v)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update value in etcd: %v", err)
	}

	return &pb.ServiceUpdateResponse{
		Result: "OK",
	}, nil
}
