// A minimal Fiber (fasthttp) service shaped for the fleet: it serves at the
// root of its own hostname, so routes mount directly on the app.
package main

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func port() string {
	if p := strings.TrimSpace(os.Getenv("PORT")); p != "" {
		return p
	}
	return "8080"
}

func newApp() *fiber.App {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(recover.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"service": "fiber-template"})
	})
	return app
}

func main() {
	addr := ":" + port()
	log.Printf("fiber-template listening on %s", addr)
	if err := newApp().Listen(addr); err != nil {
		log.Fatal(err)
	}
}
