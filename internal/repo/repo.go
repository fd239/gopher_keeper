package repo

import (
	"github.com/fd239/gopher_keeper/internal/models"
	uuid "github.com/satori/go.uuid"
)

type UsersRepo interface {
	CreateUser(user *models.User) error
	GetUserByLogin(login string) (*models.User, error)
}

type UsersDataRepo interface {
	SaveText(text *models.DataText, userId uuid.UUID) (uuid.UUID, error)
	SaveCard(card *models.DataCard, userId uuid.UUID) (uuid.UUID, error)
	SaveFile(file []byte, userId uuid.UUID) (uuid.UUID, error)
}
