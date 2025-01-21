package handlers

import (
	"errors"
	"pos-login/database/model"
	"pos-login/middleware"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var success model.Credential
	var failed model.Failed
	var userInfo model.Login
	err := c.BodyParser(&userInfo)
	if err != nil {
		failed.Message = err.Error()
		failed.Error = errors.New("bad request").Error()
		failed.StatusCode = 400

		return c.JSON(fiber.Map{
			"success":    false,
			"message":    err.Error(),
			"statusCode": 400,
			"data":       failed,
		})
	}

	accessToken, refreshToken, err := middleware.AuthenticateUser(userInfo.UserName, userInfo.Password)

	if err != nil {
		failed.Message = err.Error()
		failed.Error = errors.New("unauthorized").Error()
		failed.StatusCode = 401

		return c.JSON(fiber.Map{
			"success":    false,
			"message":    err.Error(),
			"statusCode": 401,
			"data":       failed,
		})
	}

	success.AccessToken = accessToken
	success.RefreshToken = refreshToken

	return c.JSON(fiber.Map{
		"success":    true,
		"message":    "Login Successfully",
		"statusCode": 200,
		"data":       success,
	})

}

func Refresh(c *fiber.Ctx) error {
	var data model.Credential
	header := c.Get("Authorization")
	refreshToken := strings.Replace(header, "Bearer ", "", 1)

	if refreshToken == "" {
		return c.JSON(fiber.Map{
			"success":    false,
			"message":    "No Refresh token given",
			"statusCode": 400,
		})
	}

	newAccessToken, err := middleware.RefreshAccessToken(refreshToken)
	if err != nil {
		return err
	}

	data.AccessToken = newAccessToken
	data.RefreshToken = refreshToken
	return c.JSON(fiber.Map{
		"success":    true,
		"message":    "New Access Token Generated",
		"statusCode": 200,
		"data":       data,
	})
}
