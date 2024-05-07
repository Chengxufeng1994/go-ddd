package entity

import "time"

type Permission struct {
	ID        uint
	Name      string
	Slug      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
