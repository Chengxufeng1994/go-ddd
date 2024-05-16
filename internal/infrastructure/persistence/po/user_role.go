package po

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	UserID uint `gorm:"column:user_id"`
	RoleID uint `gorm:"column:role_id"`
}

func (ur *UserRole) TableName() string {
	return fmt.Sprintf("%s.%s", SCHEMA_PREFIX, "user_roles")
}
