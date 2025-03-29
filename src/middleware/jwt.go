package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"API-Go/utils"
)

func JwtMiddleware() fiber.Handler {
	secret := utils.GetEnv("JWT_SECRET", "default_secret")
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token invalido o expirado.",
			})
		},
	})
}