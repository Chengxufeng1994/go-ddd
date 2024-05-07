package po

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Active         bool   `gorm:"default:true"`
	Email          string `gorm:"type:varchar(320);unique;not null"`
	HashedPassword string `gorm:"type:varchar(255);not null"`
	Age            int    `gorm:"type:int"`
	FirstName      string `gorm:"type:varchar(50);not null"`
	LastName       string `gorm:"type:varchar(50);not null"`
	// Address        string    `gorm:"type:text;not null"`
	// PhoneNumber    string    `gorm:"type:varchar(20);unique;not null"`
	Accounts []Account `gorm:"foreignKey:UserID;references:ID"`
}
