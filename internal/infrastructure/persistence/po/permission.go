package po

import "gorm.io/gorm"

// Permission represents the database model of permissions
type Permission struct {
	gorm.Model
	Name  string // The permission name
	Slug  string // String based unique identifier of the permission, (use hyphen seperated permission name '-', instead of space)
	Roles []Role `gorm:"many2many:role_permissions"`
}

// TableName sets the table name
func (p Permission) TableName() string {
	return "permissions"
}
