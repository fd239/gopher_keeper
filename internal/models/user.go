package models

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

// User base user struct
type User struct {
	Id       uuid.UUID
	Name     string `db:"login"`
	Password string `db:"password"`
}

// NewUser returns a new user
func NewUser(name, password string) (*User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     name,
		Password: string(pass),
	}

	return user, nil
}

// CheckPassword check password hash
func (user *User) CheckPassword(password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return
}
