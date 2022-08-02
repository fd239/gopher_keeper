package postgres

import (
	"fmt"
	"github.com/fd239/gopher_keeper/config"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Client struct {
	Conn *sqlx.DB
}

func NewPostgresClient(cfg *config.Config) (*Client, error) {
	conn, err := sqlx.Open("pgx", getDbURL(cfg))
	conn.SetMaxOpenConns(cfg.PostgresSQL.MaxOpenConns)
	conn.SetMaxIdleConns(cfg.PostgresSQL.MaxIdleConns)
	conn.SetConnMaxLifetime(cfg.PostgresSQL.ConnMaxLifetime)
	return &Client{
		Conn: conn,
	}, err
}

func getDbURL(cfg *config.Config) string {
	return fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable&search_path=public&application_name=gopher_keeper", cfg.PostgresSQL.User, cfg.PostgresSQL.Password, cfg.PostgresSQL.Address, cfg.PostgresSQL.DatabaseName)
}
