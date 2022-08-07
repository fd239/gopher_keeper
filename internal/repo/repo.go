package repo

import (
	"bytes"
	"context"
	"github.com/fd239/gopher_keeper/internal/models"
	uuid "github.com/satori/go.uuid"
)

//UsersRepo auth and register repo
type UsersRepo interface {
	CreateUser(user *models.User) error
	GetUserByLogin(login string) (*models.User, error)
}

//UsersDataRepo text and card repo
type UsersDataRepo interface {
	SaveText(text *models.DataText, userId uuid.UUID) (uuid.UUID, error)
	SaveCard(card *models.DataCard, userId uuid.UUID) (uuid.UUID, error)
}

//UsersFilesRepo file repo
type UsersFilesRepo interface {
	// Save saves a new file to the store
	Save(ctx context.Context, fileType string, fileData bytes.Buffer) (string, error)
}
