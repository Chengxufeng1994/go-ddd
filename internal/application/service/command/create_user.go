package command

import "time"

type CreateUserCommand struct {
	Email     string
	Password  string
	Age       int
	FirstName string
	LastName  string
}

type CreateUserCommandResult struct {
	ID        uint
	Active    bool
	Email     string
	Age       int
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
