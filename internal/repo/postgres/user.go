package postgres

import (
	"github.com/fd239/gopher_keeper/internal/models"
	"github.com/fd239/gopher_keeper/internal/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) repo.UsersRepo {
	return &userRepo{db: db}
}

const createStmt = `INSERT INTO users (login, password) VALUES ($1, $2)`

// CreateUser implement save user to storage
func (r *userRepo) CreateUser(user *models.User) error {
	_, err := r.db.Exec(
		createStmt,
		user.Name,
		user.Password,
	)

	if err != nil {
		return err
	}

	return nil
}

const selectByLoginStmt = `SELECT id, login, password FROM users WHERE login=$1`

// GetUserByLogin find and returns user by login
func (r *userRepo) GetUserByLogin(login string) (*models.User, error) {
	var user models.User
	err := r.db.Get(&user, selectByLoginStmt, login)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
