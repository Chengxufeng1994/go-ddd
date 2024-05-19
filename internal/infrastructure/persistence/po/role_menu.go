package po

import (
	"fmt"

	"gorm.io/gorm"
)

type RoleMenu struct {
	gorm.Model
	RoleID uint `gorm:"uniqueIndex:idx_role_id_menu_id;column:role_id"`
	MenuID uint `gorm:"uniqueIndex:idx_role_id_menu_id;column:menu_id"`
}

func (rm *RoleMenu) TableName() string {
	return fmt.Sprintf("%s.%s", SCHEMA_PREFIX, "role_menus")
}
