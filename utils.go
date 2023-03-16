package main

import "fmt"

type Converter struct{}

type Date struct{}

type FileStorage struct{}

type General struct{}

type Http struct{}

type Jwt struct {
	hmacTokenSecret   []byte
	hmacRefreshSecret []byte
	coreApiKey        string
}

type Otp struct{}

type Pagination struct{}

type Redis struct{}

type Util struct {
	Converter   *Converter
	Data        *Date
	FileStorage *FileStorage
	General     *General
	Http        *Http
	Jwt         *Jwt
	Otp         *Otp
	Pagination  *Pagination
	Redis       *Redis
}

func SetupJwt(hmacTokenSecret, hmacRefreshSecret string) *Jwt {
	fmt.Println("Setting up JWT")
	return &Jwt{
		hmacTokenSecret:   []byte(hmacTokenSecret),
		hmacRefreshSecret: []byte(hmacRefreshSecret),
	}
}

func Init(jwt *Jwt) *Util {
	fmt.Println("Setting up Utils")
	return &Util{
		Jwt: jwt,
	}
}
