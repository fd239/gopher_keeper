package repo

import "github.com/fd239/gopher_keeper/internal/models"

type UsersRepo interface {
	CreateUser(user *models.User) error
	GetUserByLogin(login string) (*models.User, error)
}

type UsersDataRepo interface {
	SaveText(text *models.DataText, userId uint) error
	SaveCard(card *models.DataCard, userId uint) error
	SaveFile(file []byte, userId uint) error
}
