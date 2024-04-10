package entity

import (
	"time"

	"github.com/Chengxufeng1994/go-ddd/internal/domain/valueobject"
)

type Customer struct {
	ID             uint
	Active         bool
	Email          *valueobject.Email
	HashedPassword string
	CustomerInfo   *valueobject.CustomerInfo
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Customers []Customer
