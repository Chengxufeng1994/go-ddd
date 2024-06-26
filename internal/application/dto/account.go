package dto

import "time"

type AccountCreationRequest struct {
	UserID   uint   `json:"user_id"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

type AccountCreationResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Amount    int64     `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	Amount    int64     `json:"amount"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
