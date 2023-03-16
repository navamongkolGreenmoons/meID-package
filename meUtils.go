package meUtils

type Converter struct{}

type Date struct{}

type Decrypt struct{}

type Http struct{}

type Jwt struct {
	hmacTokenSecret   []byte
	hmacRefreshSecret []byte
	coreApiKey        string
}

type Otp struct{}

type Pagination struct{}

type String struct{}

type MeUtil struct {
	Converter  *Converter
	Date       *Date
	Decrypt    *Decrypt
	Http       *Http
	Jwt        *Jwt
	Otp        *Otp
	Pagination *Pagination
	String     *String
}

func SetupJwt(hmacTokenSecret, hmacRefreshSecret string) *Jwt {
	return &Jwt{
		hmacTokenSecret:   []byte(hmacTokenSecret),
		hmacRefreshSecret: []byte(hmacRefreshSecret),
	}
}

func Init(jwt *Jwt) *MeUtil {
	return &MeUtil{
		Jwt: jwt,
	}
}
