package models

import (
	"golang.org/x/crypto/bcrypt"
)

// User base user struct
type User struct {
	Id       uint
	Name     string
	Password string
	Role     string
}

// NewUser returns a new user
func NewUser(name string, password string, role string) (*User, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return nil, err
	}

	user := &User{
		Name:     name,
		Password: string(pass),
		Role:     role,
	}

	return user, nil
}

// CheckPassword check password hash
func (user *User) CheckPassword(password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return
}
