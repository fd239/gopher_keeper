package postgres

import (
	"github.com/fd239/gopher_keeper/internal/repo"
	"github.com/jmoiron/sqlx"
)

type userDataRepo struct {
	db *sqlx.DB
}

func NewUserDataRepo(db *sqlx.DB) repo.UsersRepo {
	return &userRepo{db: db}
}
