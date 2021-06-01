package model

import (
	"errors"
)

type User struct {
	Id   int
	Name string
	Role string
}

type Users []User

// Exists checks if a user with  given Id exists
func (u *Users) Exists(id int) bool {
	for _, user := range *u {
		if user.Id == id {
			return true
		}
	}

	return false
}

// FindByName returns the user with the given name, or returns an error
func (u *Users) FindByName(name string) (User, error) {
	for _, user := range *u {
		if user.Name == name {
			return user, nil
		}
	}

	return User{}, errors.New("user not found")
}
