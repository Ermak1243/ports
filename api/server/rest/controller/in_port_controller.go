package controller

import (
	"net/http"
	"ports/internal/services/ports"

	"github.com/gofiber/fiber/v2"
)

// inPortController handles requests related to input ports.
type inPortController struct {
	portService ports.Port
}

// NewInPortController creates a new instance of inPortController.
func NewInPortController(portService ports.Port) *inPortController {
	return &inPortController{
		portService: portService,
	}
}

// Read handles the HTTP GET request to read a value from an input port.
// @Summary Read value from an input port
// @Description Retrieve a value from the specified input port by number
// @Tags Input Ports
// @Accept json
// @Produce json
// @Param number query int true "Port number" minimum(1)
// @Success 200 {object} object{value=string} "Successful response"
// @Failure 400 {object} object{error=string} "Invalid port number"
// @Failure 500 {object} object{error=string} "Internal server error"
// @Router /api/ports/read [get]
func (ic *inPortController) Read(c *fiber.Ctx) error {
	number := c.QueryInt("number")
	if number < 1 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "port number must be greater than 0",
		})
	}

	// Retrieve the input port using the port service.
	port, err := ic.portService.GetInPort(number)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Attempt to read a value from the result channel of the port.
	val, ok := <-port.ResultChan()
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "something went wrong",
		})
	}

	return c.JSON(fiber.Map{
		"value": val,
	})
}
