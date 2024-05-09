package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `gorm:"type:varchar(50);unique;not null"` // The name of the role
	Slug        string       `gorm:"type:varchar(50);not null"`        // String based unique identifier of the role, (use hyphen seperated role name '-', instead of space)
	Users       []User       `gorm:"many2many:user_roles"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

// TableName sets the table name
func (r Role) TableName() string {
	return "roles"
}
