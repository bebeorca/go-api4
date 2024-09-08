package request

import "time"

type UserCreateRequest struct {
	Nama      string `json:"nama"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	CreatedAt time.Time
}
