package request

type TokenUpdateRequest struct{
	IsRedeemed bool `json:"is_redeemed" gorm:"default: false"`
}