package route

import (
	"ports/internal/services/ports"

	"github.com/gofiber/fiber/v2"
)

// @title		Ports
// @version	1.0
func Setup(
	fiber *fiber.App,
	portService ports.Port,
) {
	//Groups
	api := fiber.Group("/api")
	ports := api.Group("/ports")
	docs := api.Group("/docs")

	NewDocsRouter(docs)
	NewInPortsRoute(ports, portService)
	NewOutPortsRoute(ports, portService)
}
