package dto

type RefreshTokenRequestDto struct {
	RefreshToken string `binding:"required" json:"refreshToken"`
	DeviceToken  string `binding:"required" json:"deviceToken"`
	DeviceType   string `binding:"required" json:"deviceType"`
}
