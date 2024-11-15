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

func TestReadInPortController(t *testing.T) {
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
			name:           "Bad Request. Number < 1",
			portNumber:     0,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Bad Request. No port with exact number",
			portNumber:     500,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		tc := test

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			portService := ports.NewPort(4, 7)
			inPortController := controller.NewInPortController(portService)

			app := fiber.New()
			app.Get("/api/ports/read", inPortController.Read)

			req := httptest.NewRequest("GET", fmt.Sprintf("/api/ports/read?number=%d", tc.portNumber), nil)

			resp, _ := app.Test(req, -1)

			assert.Equal(t, tc.expectedStatus, resp.StatusCode)
		})
	}
}
