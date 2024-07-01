package request


type UserCreateRequest struct{
	Nama      string         `gorm:"type:varchar(255)" json:"nama"`
	Address   string         `gorm:"type:varchar(255)" json:"address"`
	Phone     string         `gorm:"type:varchar(255)" json:"phone"`
	Email     string         `gorm:"type:varchar(255)" json:"email"`
}