package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"ports/api/server/rest/controller"
	"ports/internal/services/ports"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestWriteOutPortController(t *testing.T) {
	tests := []struct {
		name           string
		portNumber     int
		expectedStatus int
	}{
		{
			name:           "Success",
			portNumber:     1,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Bad Request. Port number lower then 1",
			portNumber:     0,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Bad Request. Port number not exists",
			portNumber:     9999,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			portService := ports.NewPort(4, 7)
			inPortController := controller.NewOutPortController(portService)

			app := fiber.New()
			app.Post("/api/ports/write", inPortController.Write)

			req := httptest.NewRequest("POST", fmt.Sprintf("/api/ports/write?port_number=%d&transaction_id=20", tc.portNumber), nil)

			resp, _ := app.Test(req, -1)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode)
		})
	}
}
