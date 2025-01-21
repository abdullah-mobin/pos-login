package config

import "os"

var (
	DBUser          string
	DBPassword      string
	DBHost          string
	DBPort          string
	DBName          string
	JWTSecret       []byte
	AccessTokenExp  string
	RefreshTokenExp string
	Port            string
)

func LoadEnv() {
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBName = os.Getenv("DB_NAME")
	JWTSecret = []byte(os.Getenv("JWT_SECRET"))
	AccessTokenExp = os.Getenv("ACCESS_TOKEN_EXP")
	RefreshTokenExp = os.Getenv("REFRESH_TOKEN_EXP")
	Port = os.Getenv("LISTENING_PORT")
}
