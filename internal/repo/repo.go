package repo

import (
	"bytes"
	"context"
	"github.com/fd239/gopher_keeper/internal/models"
	uuid "github.com/satori/go.uuid"
)

//UsersRepo auth and register repo
type UsersRepo interface {
	//CreateUser creates user
	CreateUser(user *models.User) error
	//GetUserByLogin gets user by user login
	GetUserByLogin(login string) (*models.User, error)
}

//UsersDataRepo text and card repo
type UsersDataRepo interface {
	//SaveText saves text data to storage
	SaveText(text *models.DataText, userId uuid.UUID) (uuid.UUID, error)
	//SaveCard saves card data to storage
	SaveCard(card *models.DataCard, userId uuid.UUID) (uuid.UUID, error)
	//GetText gets text data from storage
	GetText(id uuid.UUID) (*models.DataText, error)
	//GetCard gets card data from storage
	GetCard(id uuid.UUID) (*models.DataCard, error)
}

//UsersFilesRepo file repo
type UsersFilesRepo interface {
	//Save saves file to storage
	Save(ctx context.Context, fileType string, fileData bytes.Buffer) (string, error)
}
