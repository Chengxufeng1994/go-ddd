package po

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Active         bool      `gorm:"default:true"`
	Email          string    `gorm:"type:varchar(320);unique;not null"`
	HashedPassword string    `gorm:"type:varchar(255);not null"`
	Age            int       `gorm:"type:int"`
	FirstName      string    `gorm:"type:varchar(50);not null"`
	LastName       string    `gorm:"type:varchar(50);not null"`
	Enable         bool      `gorm:"type:boolean;default:true"`
	Accounts       []Account `gorm:"foreignKey:UserID;references:ID"`
	RoleID         uint      `gorm:"comment:使用者角色ID"`
	Role           Role      `gorm:"foreignKey:RoleID;references:ID;comment:使用者角色"`
	Roles          []Role    `gorm:"many2many:user_roles"`
}

// TableName sets the table name
func (u User) TableName() string {
	return fmt.Sprintf("%s.%s", SCHEMA_PREFIX, "users")
}
