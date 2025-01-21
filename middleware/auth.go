package middleware

import (
	"context"
	"errors"
	"pos-login/config"
	"pos-login/database"
	"pos-login/database/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateUser(userName string, password string) (string, string, error) {
	var user model.User
	query := "SELECT id, username, password FROM users WHERE username=$1"
	err := database.DB.QueryRow(context.Background(), query, userName).Scan(&user.ID, &user.UserName, &user.Password)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	if user.Password != password {
		return "", "", errors.New("invalid password")
	}

	accessToken, err := GenerateAccessToken(&user)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(&user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil

}

func RefreshAccessToken(refreshToken string) (string, error) {

	token, err := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return config.JWTSecret, nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	userID := claims.Issuer
	var user model.User
	query := "SELECT id, username FROM users WHERE id=$1"
	if err := database.DB.QueryRow(context.Background(), query, userID).Scan(&user.ID, &user.UserName); err != nil {
		return "", errors.New("user not found")
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		return "", errors.New("expired refresh token")

	}

	newAccessToken, err := GenerateAccessToken(&user)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil
}
