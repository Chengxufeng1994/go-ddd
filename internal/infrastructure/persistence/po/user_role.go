package po

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserID uint `gorm:"column:user_id"`
	RoleID uint `gorm:"column:role_id"`
}

func (ur *UserRole) TableName() string {
	return "user_roles"
}
