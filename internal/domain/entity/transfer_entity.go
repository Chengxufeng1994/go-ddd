package entity

import (
	"time"
)

type Transfer struct {
	ID            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Transfers []Transfer
