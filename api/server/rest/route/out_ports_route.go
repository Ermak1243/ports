package route

import (
	"ports/api/server/rest/controller"
	"ports/internal/services/ports"

	"github.com/gofiber/fiber/v2"
)

func NewOutPortsRoute(
	group fiber.Router,
	portService ports.Port,
) {
	outPortController := controller.NewOutPortController(portService)
	group.Post("/write", outPortController.Write)
}
