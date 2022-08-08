package server

import (
	"bytes"
	"context"
	"fmt"
	"github.com/fd239/gopher_keeper/internal/app/jwt"
	"github.com/fd239/gopher_keeper/internal/models"
	"github.com/fd239/gopher_keeper/internal/repo"
	"github.com/fd239/gopher_keeper/pkg/pb"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

// maxFileSize set max file upload size
const maxFileSize = 1 << 20

//userDataServer for user data handle
type userDataServer struct {
	pb.UnimplementedUserDataServiceServer
	db        repo.UsersDataRepo
	jwt       *jwt.JWTManager
	log       *zap.SugaredLogger
	fileStore repo.UsersFilesRepo
}

//NewUserDataServer returns new user file data handler
func NewUserDataServer(db repo.UsersDataRepo, jwt *jwt.JWTManager, log *zap.SugaredLogger, fileStore repo.UsersFilesRepo) *userDataServer {
	return &userDataServer{db: db, jwt: jwt, log: log, fileStore: fileStore}
}

// SaveText implement save text type data
func (r *userDataServer) SaveText(ctx context.Context, req *pb.SaveTextRequest) (*pb.SaveTextResponse, error) {
	textData := models.NewDataText(
		req.GetText().Text,
		req.GetText().GetMeta(),
	)

	userId := ctx.Value(UserIDKey{})
	userUuid, err := uuid.FromString(userId.(string))

	if err != nil {
		return nil, err
	}

	textId, err := r.db.SaveText(textData, userUuid)
	if err != nil {
		r.log.Errorf("PG save text error: %v", err)
		return nil, err
	}

	res := &pb.SaveTextResponse{Id: fmt.Sprintf("%v", textId.String())}

	return res, nil
}

// GetText implements return data text proto message
func (r *userDataServer) GetText(_ context.Context, req *pb.GetTextRequest) (*pb.GetTextResponse, error) {
	textUuid, err := uuid.FromString(req.GetId())

	if err != nil {
		return nil, err
	}

	textData, err := r.db.GetText(textUuid)
	if err != nil {
		r.log.Errorf("PG get text error: %v", err)
		return nil, err
	}

	res := &pb.GetTextResponse{Text: textData.ToProto()}

	return res, nil
}

// SaveCard implement save card type data
func (r *userDataServer) SaveCard(ctx context.Context, req *pb.SaveCardRequest) (*pb.SaveCardResponse, error) {
	cardData := models.NewDataCard(
		req.GetCard().GetNumber(),
		req.GetCard().GetMeta(),
	)

	userId := ctx.Value(UserIDKey{})
	userUuid, err := uuid.FromString(userId.(string))

	if err != nil {
		r.log.Errorf("PG save card error: %v", err)
		return nil, err
	}

	cardId, err := r.db.SaveCard(cardData, userUuid)
	if err != nil {
		return nil, err
	}

	res := &pb.SaveCardResponse{Id: fmt.Sprintf("%v", cardId.String())}

	return res, nil
}

// GetCard implements return data card proto message
func (r *userDataServer) GetCard(_ context.Context, req *pb.GetCardRequest) (*pb.GetCardResponse, error) {
	cardUuid, err := uuid.FromString(req.GetId())

	if err != nil {
		return nil, err
	}

	cardData, err := r.db.GetCard(cardUuid)
	if err != nil {
		r.log.Errorf("PG get card error: %v", err)
		return nil, err
	}

	res := &pb.GetCardResponse{Card: cardData.ToProto()}

	return res, nil
}

// SaveFile implements save file type data
func (r *userDataServer) SaveFile(stream pb.UserDataService_SaveFileServer) error {
	req, err := stream.Recv()
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot receive image info")
	}

	fileId := req.GetInfo().GetId()
	fileType := req.GetInfo().GetType()
	r.log.Infof("receive an upload-image request for file %s with image type %s", fileId, fileType)

	fileData := bytes.Buffer{}
	fileSize := 0

	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		r.log.Info("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			r.log.Info("no more data")
			break
		}
		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		r.log.Info("received a chunk with size: %d", size)

		fileSize += size
		if fileSize > maxFileSize {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", fileSize, maxFileSize)
		}

		// write slowly
		// time.Sleep(time.Second)

		_, err = fileData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}
	}

	imageID, err := r.fileStore.Save(stream.Context(), fileType, fileData)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot save image to the store: %v", err)
	}

	res := &pb.FileResponse{
		Id:   imageID,
		Size: uint32(fileSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot send response: %v", err)
	}

	r.log.Infof("saved image with id: %s, size: %d", imageID, fileSize)

	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}
