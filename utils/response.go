package utils

import "github.com/gofiber/fiber/v2"

func SendSuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(fiber.Map{
		"success":    true,
		"message":    message,
		"statusCode": 200,
		"data":       data,
	})
}

func SendFailedResponse(c *fiber.Ctx, message string, code int) error {
	return c.JSON(fiber.Map{
		"success":    false,
		"message":    message,
		"statusCode": code,
	})
}
