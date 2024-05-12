package dto

import "time"

type UserCreationRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Age       int    `json:"age"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserCreationResponse struct {
	ID        uint      `json:"id"`
	Active    bool      `json:"active"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID        uint      `json:"id"`
	Active    bool      `json:"active"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserQueryParams struct {
	CurrentPage int    `json:"current_page" form:"current_page,default=1"`
	PageSize    int    `json:"page_size" form:"page_size,default=10"`
	SortBy      string `json:"sort_by" form:"sort_by,default=created_at"`
	OrderBy     string `json:"order_by" form:"order_by,default=asc"`
	Email       string `json:"email" form:"email"`
}
