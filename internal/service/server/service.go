package server

import (
	"context"
	"github.com/fd239/gopher_keeper/config"
	"github.com/fd239/gopher_keeper/internal/app/jwt"
	"github.com/fd239/gopher_keeper/internal/app/server"
	"github.com/fd239/gopher_keeper/internal/repo/postgres"
	"github.com/fd239/gopher_keeper/pkg/crypt"
	"github.com/fd239/gopher_keeper/pkg/pb"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	certFile        = "ssl/server.crt"
	keyFile         = "ssl/server.pem"
	maxHeaderBytes  = 1 << 20
	gzipLevel       = 5
	stackSize       = 1 << 10 // 1 KB
	csrfTokenHeader = "X-CSRF-Token"
	bodyLimit       = "2M"
	jwtSecret       = "test"
	jwtTimer        = 1 * time.Hour
)

type service struct {
	log   *zap.SugaredLogger
	cfg   *config.Config
	db    *sqlx.DB
	crypt *crypt.CipherCrypt
}

// NewService constructor
func NewService(
	log *zap.SugaredLogger,
	cfg *config.Config,
	db *sqlx.DB,
	crypt *crypt.CipherCrypt,
) *service {
	return &service{log: log, cfg: cfg, db: db, crypt: crypt}
}

// Run start application
func (s *service) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//emailHandlers := emailsV1.NewEmailHandlers(v1.Group("/email"), emailUC, s.log, validate)
	//emailHandlers.MapRoutes()

	l, err := net.Listen("tcp", s.cfg.GRPC.Port)
	if err != nil {
		return errors.Wrap(err, "net.Listen")
	}
	defer l.Close()

	//cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	//if err != nil {
	//	s.log.Fatalf("failed to load key pair: %s", err)
	//}

	jwtManager := jwt.NewJWTManager(jwtSecret, jwtTimer)
	authInterceptor := server.NewAuthInterceptor(jwtManager, publicMethods())

	grpcServer := grpc.NewServer(
		//grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: s.cfg.GRPC.MaxConnectionIdle * time.Minute,
			Timeout:           s.cfg.GRPC.Timeout * time.Second,
			MaxConnectionAge:  s.cfg.GRPC.MaxConnectionAge * time.Minute,
			Time:              s.cfg.GRPC.Timeout * time.Minute,
		}),
		grpc.UnaryInterceptor(authInterceptor.Unary()),
	)

	//Auth
	userRepo := postgres.NewUserRepo(s.db)
	authServer := server.NewAuthServer(userRepo, jwtManager)

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	//User data
	userDataRepo := postgres.NewUserDataRepo(s.db, s.crypt)
	userDataServer := server.NewUserDataServer(userDataRepo, jwtManager)
	pb.RegisterUserDataServiceServer(grpcServer, userDataServer)

	grpc_prometheus.Register(grpcServer)

	s.log.Infof("GRPC Server is listening on port: %s", s.cfg.GRPC.Port)
	s.log.Fatal(grpcServer.Serve(l))

	if s.cfg.HTTP.Development {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		s.log.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		s.log.Errorf("ctx.Done: %v", done)
	}

	grpcServer.GracefulStop()
	s.log.Info("Graceful shutdown")

	return nil
}

func publicMethods() map[string]bool {
	return map[string]bool{
		"/api.AuthService/Register": true,
		"/api.AuthService/Login":    true,
	}
}
