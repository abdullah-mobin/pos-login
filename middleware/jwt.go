package middleware

import (
	"errors"
	"fmt"
	"pos-login/config"
	"pos-login/database/model"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(c *fiber.Ctx) error {

	var failed model.Failed
	tokenStr := c.Get("Authorization")

	tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

	if tokenStr == "" {
		failed.Message = errors.New("token not given").Error()
		failed.Error = errors.New("bad request").Error()
		failed.StatusCode = 400
		return c.JSON(fiber.Map{
			"success":    false,
			"message":    failed.Message,
			"statusCode": 400,
			"data":       failed,
		})

	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		failed.Message = errors.New("can not valid token").Error()
		failed.Error = errors.New("unauthorized").Error()
		failed.StatusCode = 401
		return c.JSON(fiber.Map{
			"success":    false,
			"message":    failed.Message,
			"statusCode": 401,
			"data":       failed,
		})
	}

	return c.Next()
}

func GenerateAccessToken(user *model.User) (string, error) {
	accessTokenExp, err := time.ParseDuration(config.AccessTokenExp)
	if err != nil {
		return "", err
	}
	expTime := time.Now().Add(time.Duration(accessTokenExp) * time.Minute)

	claims := &jwt.RegisteredClaims{
		Issuer:    fmt.Sprintf("%d", user.ID),
		ExpiresAt: jwt.NewNumericDate(expTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.JWTSecret)
}

func GenerateRefreshToken(user *model.User) (string, error) {
	refreshTokenExp, err := time.ParseDuration(config.RefreshTokenExp)
	if err != nil {
		return "", err
	}
	expTime := time.Now().Add(time.Duration(refreshTokenExp) * time.Minute)

	claims := &jwt.RegisteredClaims{
		Issuer:    fmt.Sprintf("%d", user.ID),
		ExpiresAt: jwt.NewNumericDate(expTime),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(config.JWTSecret)
}
