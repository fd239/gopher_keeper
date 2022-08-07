package server

import (
	"context"
	"github.com/fd239/gopher_keeper/internal/app/jwt"
	"github.com/fd239/gopher_keeper/internal/models"
	"github.com/fd239/gopher_keeper/internal/repo"
	"github.com/fd239/gopher_keeper/pkg/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// authServer service for user registration
type authServer struct {
	pb.UnimplementedAuthServiceServer
	db         repo.UsersRepo
	jwtManager *jwt.JWTManager
}

// New returns a new auth server
func NewAuthServer(db repo.UsersRepo, jwtManager *jwt.JWTManager) *authServer {
	return &authServer{db: db, jwtManager: jwtManager}
}

// Register registration user logic
func (s *authServer) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, err := s.db.GetUserByLogin(req.GetUsername())
	if err == nil {
		return nil, status.Errorf(codes.AlreadyExists, "user exist error: %v", err)
	}

	// Create new user
	user, err := models.NewUser(req.GetUsername(), req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "user model creation error: %v", err)
	}

	if err = s.db.CreateUser(user); err != nil {
		return nil, status.Errorf(codes.Internal, "db user creation error: %v", err)
	}

	token, err := s.jwtManager.GenerateJWT(user.Id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.RegisterResponse{AccessToken: token}

	return res, nil
}

// Login logging in user logic
func (s *authServer) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.db.GetUserByLogin(req.GetUsername())
	if user == nil {
		return nil, status.Errorf(codes.AlreadyExists, "no user found by username: %v", req.GetUsername())
	}

	if err = user.CheckPassword(req.GetPassword()); err != nil {
		return nil, status.Errorf(codes.NotFound, "password error: %v", err)
	}

	token, err := s.jwtManager.GenerateJWT(user.Id.String())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{AccessToken: token}

	return res, nil
}
