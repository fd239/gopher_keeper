package postgres

import (
	"github.com/fd239/gopher_keeper/internal/models"
	"github.com/fd239/gopher_keeper/internal/repo"
	"github.com/jmoiron/sqlx"
)

type userDataRepo struct {
	db *sqlx.DB
}

func NewUserDataRepo(db *sqlx.DB) repo.UsersDataRepo {
	return &userDataRepo{db: db}
}

const saveTextStmt = `INSERT INTO users (text, id, user_id, meta, type) VALUES ($1, $2, $3, $4, $5)`

// CreateUser implement save user to storage
func (r *userDataRepo) SaveText(data *models.DataText, userId uint) error {
	_, err := r.db.Exec(
		saveTextStmt,
		data.Text,
		data.Id,
		data.UserId,
		data.Meta,
		data.Type,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *userDataRepo) SaveCard(card *models.DataCard, userId uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *userDataRepo) SaveFile(file []byte, userId uint) error {
	//TODO implement me
	panic("implement me")
}
