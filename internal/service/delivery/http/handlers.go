package http

import (
	"path/filepath"

	"github.com/AlexandrKobalt/go-files/config"
	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	cfg *config.Config
}

func New(cfg *config.Config) IHandlers {
	return &handlers{cfg: cfg}
}

func (h *handlers) ServeFile() fiber.Handler {
	return func(c *fiber.Ctx) error {
		uuid := c.Params("key")
		filePath := filepath.Join(h.cfg.Path, uuid)

		return c.SendFile(filePath)
	}
}
