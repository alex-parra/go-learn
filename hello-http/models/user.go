package models

import (
	"fmt"
)

// User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

var (
	users  []*User
	nextID = 0
)

// GetUsers to get all users
func GetUsers() []*User {
	return users
}

// AddUser to add a user
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, fmt.Errorf("New user can't have an ID")
	}

	nextID++
	u.ID = nextID
	users = append(users, &u)
	return u, nil
}

// GetUserByID to get a user by ID
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser to update a user
func UpdateUser(u User) (User, error) {
	for i, user := range users {
		if user.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("Unknown user")
}

// RemoveUser to remove user
func RemoveUser(u User) error {
	for i, user := range users {
		if user.ID == u.ID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User not found")
}
