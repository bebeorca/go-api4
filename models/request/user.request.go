package request

type UserCreateRequest struct {
	Nama      string `json:"nama"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
