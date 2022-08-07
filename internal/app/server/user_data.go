package server

import (
	"context"
	"fmt"
	"github.com/fd239/gopher_keeper/internal/app/jwt"
	"github.com/fd239/gopher_keeper/internal/models"
	"github.com/fd239/gopher_keeper/internal/repo"
	"github.com/fd239/gopher_keeper/pkg/pb"
	uuid "github.com/satori/go.uuid"
)

type userDataServer struct {
	pb.UnimplementedUserDataServiceServer
	db  repo.UsersDataRepo
	jwt *jwt.JWTManager
}

func NewUserDataServer(db repo.UsersDataRepo, jwt *jwt.JWTManager) *userDataServer {
	return &userDataServer{db: db, jwt: jwt}
}

// SaveText implement save text type data
func (r *userDataServer) SaveText(ctx context.Context, req *pb.TextRequest) (*pb.TextResponse, error) {
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
		return nil, err
	}

	res := &pb.TextResponse{Id: fmt.Sprintf("%v", textId.String())}

	return res, nil
}
