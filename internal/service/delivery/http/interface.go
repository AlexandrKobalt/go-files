package http

import "github.com/gofiber/fiber/v2"

type IHandlers interface {
	ServeFile() fiber.Handler
}
