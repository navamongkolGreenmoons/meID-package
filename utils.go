package utils

type Converter struct{}

type Date struct{}

type Decrypt struct{}

type General struct{}

type Http struct{}

type Jwt struct {
	hmacTokenSecret   []byte
	hmacRefreshSecret []byte
	coreApiKey        string
}

type Otp struct{}

type Pagination struct{}

type Util struct {
	Converter  *Converter
	Date       *Date
	Decrypt    *Decrypt
	General    *General
	Http       *Http
	Jwt        *Jwt
	Otp        *Otp
	Pagination *Pagination
}

func SetupJwt(hmacTokenSecret, hmacRefreshSecret string) *Jwt {
	return &Jwt{
		hmacTokenSecret:   []byte(hmacTokenSecret),
		hmacRefreshSecret: []byte(hmacRefreshSecret),
	}
}

func Init(jwt *Jwt) *Util {
	return &Util{
		Jwt: jwt,
	}
}
