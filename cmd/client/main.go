package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/fd239/gopher_keeper/config"
	"github.com/fd239/gopher_keeper/internal/service/client"
	"github.com/fd239/gopher_keeper/pkg/logger"
	"github.com/fd239/gopher_keeper/pkg/pb"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"time"
)

const (
	username      = "test"
	dummyPassword = "hello123"
)

func main() {
	cfg := config.ParseConfig(".env.client")

	appLogger := logger.NewLogger(cfg)
	appLogger.Info("Starting gopher keeper client")

	conn, _, userData := client.NewUserClientWithConnection(cfg.GRPC.Port, true, &pb.LoginRequest{
		Username: "dummy-1659893821255105800",
		Password: "hello123",
	})
	defer conn.Close()

	//registerResp, err := uc.Register(context.Background(), getDummyRegistrationRequest(true))
	//if err != nil {
	//	appLogger.Errorf("register request error: %v", err)
	//}
	//
	//b, _ := json.MarshalIndent(registerResp, "", "\t")
	//appLogger.Info("Register result:")
	//appLogger.Info(string(b))

	saveResp, err := userData.SaveText(context.Background(), &pb.TextRequest{Text: &pb.DataText{
		Text: "test",
		Meta: "test",
	}})

	if err != nil {
		appLogger.Errorf("save text error: %v", err)
	}

	b, _ := json.MarshalIndent(saveResp, "", "\t")
	appLogger.Info("save text result:")
	appLogger.Info(string(b))
	////resp, err := client.Register(ctx, &pb.RegisterRequest{
	////	Username: "test",
	////	Password: "test",
	////})
	////
	////if err != nil {
	////	appLogger.Errorf("register request error: %v", err)
	////	return
	////}
	////
	////bytes, _ := json.MarshalIndent(resp, "", "\t")
	////appLogger.Info("Register result:")
	////appLogger.Info(string(bytes))
	//
	//loginResp, err := clientAuth.Login(ctx, &pb.LoginRequest{
	//	Username: "test",
	//	Password: "test",
	//})
	//
	//if err != nil {
	//	appLogger.Errorf("login request error: %v", err)
	//	return
	//}
	//
	//bytes, _ := json.MarshalIndent(loginResp, "", "\t")
	//appLogger.Info("Login result:")
	//appLogger.Info(string(bytes))
	//
	//clientUserData := pb.NewUserDataServiceClient(conn)
	//clientUserData.SaveText(ctx, &pb.TextRequest{Text: &pb.DataText{
	//	Text: "test",
	//	Meta: "test",
	//}})

}

func UploadImage(userDataClient pb.UserDataServiceClient, appLogger *zap.SugaredLogger, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		appLogger.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := userDataClient.UploadImage(ctx)
	if err != nil {
		appLogger.Fatalf("cannot upload image: %v", err)
	}

	req := &pb.FileRequest{
		Data: &pb.FileRequest_Info{
			Info: &pb.FileInfo{
				Id:   "123",
				Type: filepath.Ext(imagePath),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		appLogger.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
	}

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			appLogger.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &pb.FileRequest{
			Data: &pb.FileRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			appLogger.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		appLogger.Fatal("cannot receive response: ", err)
	}

	appLogger.Infof("image uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
}

func getDummyRegistrationRequest(existingUser bool) *pb.RegisterRequest {
	rnd := time.Now().UnixNano()
	var reqUsername string
	if existingUser {
		reqUsername = fmt.Sprintf("dummy-%d", rnd)
	} else {
		reqUsername = username
	}

	return &pb.RegisterRequest{
		Username: reqUsername,
		Password: dummyPassword,
	}
}
