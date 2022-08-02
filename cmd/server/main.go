package main

import (
	"github.com/fd239/gopher_keeper/config"
	"github.com/fd239/gopher_keeper/internal/server"
	"github.com/fd239/gopher_keeper/pkg/logger"
	"github.com/fd239/gopher_keeper/pkg/postgres"
)

// @title Gopher Keeper
// @version 1.0
// @description Gopher keeper
// @termsOfService http://swagger.io/terms/

// @contact.name Dmitry Frolov
// @contact.url https://github.com/fd239
// @contact.email fd239@bk.ru

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /api/v1
func main() {
	cfg := config.ParseConfig()

	appLogger := logger.NewLogger(cfg)
	appLogger.Info("Starting OCS service")

	pgConn, err := postgres.NewPostgresClient(cfg)
	if err != nil {
		appLogger.Fatalf("Postgres connection error: %v", err)
	}
	appLogger.Info("Postgres Connected")

	sever := server.NewServer(appLogger, cfg, pgConn.Conn)
	appLogger.Fatal(sever.Run())
}
