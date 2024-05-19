package dto

import "time"

type Menu struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	Description  string    `json:"description"`
	Path         string    `json:"path"`
	ParentID     uint      `json:"parent_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ChildrenMenu []*Menu   `json:"children_menu"`
}

type MenuCreationRequest struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Path        string `json:"path"`
	ParentID    uint   `json:"parent_id"`
}
