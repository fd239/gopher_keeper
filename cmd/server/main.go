package main

import (
	"github.com/fd239/gopher_keeper/config"
	"github.com/fd239/gopher_keeper/internal/service/server"
	"github.com/fd239/gopher_keeper/migrations"
	"github.com/fd239/gopher_keeper/pkg/crypt"
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
	cfg := config.ParseConfig(".env")

	appLogger := logger.NewLogger(cfg)
	appLogger.Info("Starting gopher keeper server")

	pgConn, err := postgres.NewPostgresClient(cfg)
	if err != nil {
		appLogger.Fatalf("Postgres connection error: %v", err)
	}
	appLogger.Info("Postgres Connected")

	crypt, err := crypt.NewCrypt(cfg)
	if err != nil {
		appLogger.Fatalf("Crypt init error: %v", err)
	}

	if version, err := migrations.Run(pgConn.Conn.DB); err != nil {
		appLogger.Fatalf("Error run PG migrations: %v", err)
	} else {
		appLogger.Infof("Successfully completed migrations (v.%d)", version)
	}

	sever := server.NewService(appLogger, cfg, pgConn.Conn, crypt)
	appLogger.Fatal(sever.Run())
}
