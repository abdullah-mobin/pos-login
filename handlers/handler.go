package handlers

import (
	"pos-login/database/model"
	"pos-login/middleware"
	"pos-login/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {

	var success model.Credential
	var userInfo model.Login
	err := c.BodyParser(&userInfo)
	if err != nil {
		return utils.SendFailedResponse(c, "bad request", 400)

	}

	accessToken, refreshToken, err := middleware.AuthenticateUser(userInfo.UserName, userInfo.Password)
	if err != nil {
		return utils.SendFailedResponse(c, "unauthorized", 401)
	}

	success.AccessToken = accessToken
	success.RefreshToken = refreshToken

	return utils.SendSuccessResponse(c, "Login Successfully", success)
}

func Refresh(c *fiber.Ctx) error {
	var data model.Credential
	header := c.Get("Authorization")
	refreshToken := strings.Replace(header, "Bearer ", "", 1)

	if refreshToken == "" {
		return utils.SendFailedResponse(c, "unauthorized", 401)
	}

	newAccessToken, err := middleware.RefreshAccessToken(refreshToken)
	if err != nil {
		return err
	}

	data.AccessToken = newAccessToken
	data.RefreshToken = refreshToken

	return utils.SendSuccessResponse(c, "New Access Token Generated", data)
}
