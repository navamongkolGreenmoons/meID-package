package utils

import (
	"fmt"
	"strings"
	"time"

	apiErrors "github.com/navamongkolgreenmoons/meID-package/dto/apiErrors"
	jwtDto "github.com/navamongkolgreenmoons/meID-package/dto/jwt"

	jwt "github.com/golang-jwt/jwt/v4"
)

func (j *Jwt) JwtGen(userId string) (*jwtDto.JwtResponseDto, *error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userId
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	t, err := token.SignedString(j.hmacTokenSecret)
	if err != nil {
		return nil, &err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = userId
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString(j.hmacRefreshSecret)
	if err != nil {
		return nil, &err
	}

	return &jwtDto.JwtResponseDto{
		AccessToken:  t,
		RefreshToken: rt,
		ExpiredIn:    1200,
	}, nil
}

func (j *Jwt) ValidateApiKey(apiKey string) bool {
	return j.coreApiKey == apiKey
}

func (j *Jwt) JwtValidate(tokenString string) (*string, *apiErrors.CustomError) {
	// 1. Check if token is nil or blank
	if tokenString == "" {
		return nil, apiErrors.Unauthorized("Unauthorized")
	}

	// 2. Replace 'Bearer ' to blank
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// 3. Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.hmacTokenSecret, nil
	})
	if err != nil {
		return nil, apiErrors.Unauthorized("Unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userId := fmt.Sprint(claims["sub"])

	if !ok || !token.Valid {
		return nil, apiErrors.Unauthorized("Unauthorized")
	}

	return &userId, nil
}
