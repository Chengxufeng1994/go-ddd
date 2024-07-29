package query

import "time"

type GetUserQuery struct {
	UserID uint
}

type GetUserQueryResult struct {
	ID        uint
	Active    bool
	Email     string
	Age       int
	FirstName string
	LastName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
