package controller

import (
	"net/http"
	"ports/internal/services/ports"

	"github.com/gofiber/fiber/v2"
)

// outPortController handles requests related to output ports.
type outPortController struct {
	portService ports.Port
}

// NewOutPortController creates a new instance of outPortController.
func NewOutPortController(portService ports.Port) *outPortController {
	return &outPortController{
		portService: portService,
	}
}

// Write handles the HTTP POST request to send a value to an output port.
// @Summary Send a transaction ID to an output port
// @Description Sends a transaction ID to the specified output port by port number
// @Tags Output Ports
// @Accept json
// @Produce json
// @Param port_number query int true "Port number" minimum(1)
// @Param transaction_id query int true "Transaction ID"
// @Success 200 {object} object{message=string} "Successful response"
// @Failure 400 {object} object{error=string} "Invalid port number or transaction ID"
// @Router /api/ports/write [post]
func (oc *outPortController) Write(c *fiber.Ctx) error {
	portNumber := c.QueryInt("port_number")
	transactionID := c.QueryInt("transaction_id")

	if portNumber < 1 || transactionID < 1 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "port number and transaction id must be greater than 0",
		})
	}

	// Retrieve the output port using the port service.
	port, err := oc.portService.GetOutPort(portNumber)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Send the transaction ID to the specified output port.
	port.Send(portNumber, transactionID)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
