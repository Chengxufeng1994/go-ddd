package po

import (
	"fmt"

	"gorm.io/gorm"
)

type Menu struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(50);unique;not null"` // The name of the role
	Slug         string  `gorm:"type:varchar(50);not null"`        // String based unique identifier of the role, (use hyphen seperated role name '-', instead of space)
	Description  string  `gorm:"type:text;not null"`
	Path         string  `gorm:"type:varchar(50);not null"`
	ParentID     uint    `gorm:"type:integer"`
	ChildrenMenu []*Menu `gorm:"-"`
}

// TableName sets the table name
func (m Menu) TableName() string {
	return fmt.Sprintf("%s.%s", SCHEMA_PREFIX, "menus")
}
