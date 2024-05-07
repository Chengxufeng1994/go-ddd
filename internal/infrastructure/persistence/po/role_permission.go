package po

import "gorm.io/gorm"

type RolePermission struct {
	gorm.Model
	RoleID       uint `gorm:"column:role_id"`
	PermissionID uint `gorm:"column:permission_id"`
}

func (rp *RolePermission) TableName() string {
	return "role_permissions"
}
