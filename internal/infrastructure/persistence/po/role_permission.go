package po

import (
	"fmt"

	"gorm.io/gorm"
)

type RolePermission struct {
	gorm.Model
	RoleID       uint `gorm:"uniqueIndex:idx_role_id_permission_id;column:role_id"`
	PermissionID uint `gorm:"uniqueIndex:idx_role_id_permission_id;column:permission_id"`
}

func (rp *RolePermission) TableName() string {
	return fmt.Sprintf("%s.%s", SCHEMA_PREFIX, "role_permissions")
}
