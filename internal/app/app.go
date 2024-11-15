package app

import (
	"log"
	"os"
	"os/signal"
	"ports/api/server/rest/route"
	"ports/internal/config"
	"ports/internal/services/ports"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// Run initializes the configuration, creates the server, and starts it.
func Run() {
	cfg := config.NewConfig("./configs/config.yml")
	ports := ports.NewPort(cfg.InPortsCount, cfg.OutPortsCount)

	// Initialize a new Fiber instance with custom settings for JSON encoding and decoding.
	fiber := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		Immutable:   true,
	})

	// Set up API routes, passing the Fiber instance and the Port object.
	route.Setup(fiber, ports)

	// Create a channel to intercept OS signals (e.g., Ctrl+C).
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Start a goroutine to handle the interrupt signal.
	go func() {
		<-c
		log.Println("Gracefully shutting down...")
		fiber.Shutdown()
	}()

	// Start the server on the specified port from the configuration.
	if err := fiber.Listen(cfg.ServerPort); err != nil {
		panic(err)
	}
}
