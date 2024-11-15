package route

import (
	"ports/api/server/rest/controller"
	"ports/internal/services/ports"

	"github.com/gofiber/fiber/v2"
)

func NewInPortsRoute(
	group fiber.Router,
	portService ports.Port,
) {
	inPortController := controller.NewInPortController(portService)
	group.Get("/read", inPortController.Read)
}
