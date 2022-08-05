package server

import (
	"context"
	"fmt"
	"github.com/fd239/gopher_keeper/internal/app/jwt"
	"github.com/fd239/gopher_keeper/internal/models"
	"github.com/fd239/gopher_keeper/internal/repo"
	"github.com/fd239/gopher_keeper/pkg/pb"
)

type userDataServer struct {
	pb.UnimplementedUserDataServiceServer
	db  repo.UsersDataRepo
	jwt *jwt.JWTPayload
}

func NewUserDataServer(db repo.UsersDataRepo, jwt *jwt.JWTPayload) *userDataServer {
	return &userDataServer{db: db, jwt: jwt}
}

// SaveText implement save text type data
func (r *userDataServer) SaveText(_ context.Context, req *pb.TextRequest) (*pb.TextResponse, error) {
	textData := models.NewDataText(
		int(req.GetText().GetId()),
		req.GetText().GetName(),
		req.GetText().GetMeta(),
	)

	err := r.db.SaveText(textData, r.jwt.UserId)
	if err != nil {
		return nil, err
	}

	res := &pb.TextResponse{Id: fmt.Sprintf("%v", textData.Id)}

	return res, nil
}
