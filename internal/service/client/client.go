package client

import (
	"context"
	"fmt"
	"github.com/fd239/gopher_keeper/pkg/pb"
	"google.golang.org/grpc"
)

//NewUserClientWithConnection dials a connection to gRPC server and initializes a new UserClient using the connection
func NewUserClientWithConnection(grpcServerAddr string, attachToken bool, loginRequest *pb.LoginRequest) (*grpc.ClientConn, pb.AuthServiceClient, pb.UserDataServiceClient) {
	var uc pb.AuthServiceClient
	var userData pb.UserDataServiceClient
	conn, _ := grpc.Dial(grpcServerAddr, grpc.WithInsecure())
	uc = pb.NewAuthServiceClient(conn)
	userData = pb.NewUserDataServiceClient(conn)

	if attachToken {
		login, err := uc.Login(context.Background(), loginRequest)
		if err != nil {
			fmt.Println(err)
			return nil, nil, nil
		}
		interceptor := AuthInterceptor{
			Token: login.AccessToken,
		}
		conn, _ = grpc.Dial(
			grpcServerAddr,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(interceptor.GetUnaryInterceptor()),
		)
		uc = pb.NewAuthServiceClient(conn)
		userData = pb.NewUserDataServiceClient(conn)
	}
	return conn, uc, userData
}
