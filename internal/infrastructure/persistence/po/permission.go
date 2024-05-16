package po

import (
	"fmt"

	"gorm.io/gorm"
)

// Permission represents the database model of permissions
type Permission struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50);unique;not null"` // The permission name
	Slug  string `gorm:"type:varchar(50);not null"`        // String based unique identifier of the permission, (use hyphen seperated permission name '-', instead of space)
	Roles []Role `gorm:"many2many:role_permissions"`
}

// TableName sets the table name
func (p Permission) TableName() string {
	return fmt.Sprintf("%s.%s", SCHEMA_PREFIX, "permissions")
}
