package entity

import (
// "time"
)

type Post struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ImagePath string `json:"file_path"`
	CreatedAt string `gorm:"type:varchar(255)" json:"created_at"`
}
