package server

import (
	"context"
	"crypto/tls"
	"github.com/fd239/gopher_keeper/config"
	"github.com/fd239/gopher_keeper/internal/middlewares"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
)

type server struct {
	log  *zap.SugaredLogger
	cfg  *config.Config
	db   *sqlx.DB
	echo *echo.Echo
}

// NewServer constructor
func NewServer(
	log *zap.SugaredLogger,
	cfg *config.Config,
	db *sqlx.DB,
) *server {
	return &server{log: log, cfg: cfg, db: db, echo: echo.New()}
}

// Run start application
func (s *server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//smtpClient := smtp.NewSmtpClient(s.cfg)
	//publisher := nats.NewPublisher(s.natsConn)
	//emailPgRepo := repository.NewEmailPGRepository(s.pgxPool)
	//emailRedisRepo := repository.NewEmailRedisRepository(s.redis)
	//emailUC := usecase.NewEmailUseCase(s.log, emailPgRepo, publisher, smtpClient, emailRedisRepo)
	//
	//im := interceptors.NewInterceptorManager(s.log, s.cfg)
	mw := middlewares.NewMiddlewareManager(s.log, s.cfg)
	//
	//validate := validator.New()

	//go func() {
	//	emailSubscriber := nats.NewEmailSubscriber(s.natsConn, s.log, emailUC, validate)
	//	emailSubscriber.Run(ctx)
	//}()

	go func() {
		s.log.Infof("Server is listening on PORT: %s", s.cfg.HTTP.Port)
		s.runHttpServer()
	}()

	metricsServer := echo.New()
	go func() {
		metricsServer.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		s.log.Infof("Metrics server is running on port: %s", s.cfg.Metrics.Port)
		if err := metricsServer.Start(s.cfg.Metrics.Port); err != nil {
			s.log.Error(err)
			cancel()
		}
	}()
	v1 := s.echo.Group("/api/v1")
	v1.Use(mw.Metrics)

	//emailHandlers := emailsV1.NewEmailHandlers(v1.Group("/email"), emailUC, s.log, validate)
	//emailHandlers.MapRoutes()

	l, err := net.Listen("tcp", s.cfg.GRPC.Port)
	if err != nil {
		return errors.Wrap(err, "net.Listen")
	}
	defer l.Close()

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		s.log.Fatalf("failed to load key pair: %s", err)
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: s.cfg.GRPC.MaxConnectionIdle * time.Minute,
			Timeout:           s.cfg.GRPC.Timeout * time.Second,
			MaxConnectionAge:  s.cfg.GRPC.MaxConnectionAge * time.Minute,
			Time:              s.cfg.GRPC.Timeout * time.Minute,
		}),
	)

	//emailGRPCService := emailGrpc.NewEmailGRPCService(emailUC, s.log, validate)
	//emailService.RegisterEmailServiceServer(grpcServer, emailGRPCService)
	//grpc_prometheus.Register(grpcServer)

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
