package entity

import (
	// "time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Nama      string         `gorm:"type:varchar(255)" json:"nama"`
	Address   string         `gorm:"type:varchar(255)" json:"address"`
	Phone     string         `gorm:"type:varchar(255)" json:"phone"`
	Email     string         `gorm:"type:varchar(255)" json:"email"`
	CreatedAt string         `gorm:"type:varchar(255)" json:"created_at"`
	UpdatedAt string         `gorm:"type:varchar(255)" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
