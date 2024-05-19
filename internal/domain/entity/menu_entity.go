package entity

import "time"

type Menu struct {
	ID           uint
	Name         string
	Slug         string
	Description  string
	Path         string
	ParentID     uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ChildrenMenu []*Menu
}
