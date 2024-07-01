package request

import "time"

type UserCreateRequest struct {
	Nama      string `json:"nama"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CreatedAt time.Time
}
