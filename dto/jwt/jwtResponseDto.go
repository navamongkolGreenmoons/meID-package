package dto

type JwtResponseDto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiredIn    int    `json:"expiredIn"`
}