package entity

import "time"

type Role struct {
	ID        uint
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
