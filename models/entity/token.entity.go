package entity

type Token struct{
	ID uint `gorm:"primaryKey" json:"id"`
	Token string `json:"token"`
	IsRedemeed bool `json:"is_redeemed" gorm:"default: false"`
}