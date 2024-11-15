package route

import (
	_ "ports/api/swagger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func NewDocsRouter(group fiber.Router) {
	group.Get("/swagger/*", swagger.HandlerDefault)
}
