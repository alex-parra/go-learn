package models

// User struct
type User struct {
	ID        int
	FirstName string
	LastName  string
	Age       int
}

var (
	users  []*User
	nextID = 1
)
