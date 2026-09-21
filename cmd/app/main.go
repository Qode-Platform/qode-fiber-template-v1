// A minimal Fiber (fasthttp) service shaped for the fleet.
package main

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// basePath returns the fleet's ingress prefix, normalised to "" or
// "/leading/no-trailing-slash". The fleet injects BASE_PATH as
// /direct/<agent>:<port> and nginx forwards it UNCHANGED, so every route must
// live under it. Empty means standalone: serve at the host root.
func basePath() string {
	raw := strings.Trim(strings.TrimSpace(os.Getenv("BASE_PATH")), "/")
	if raw == "" {
		return ""
	}
	return "/" + raw
}

func port() string {
	if p := strings.TrimSpace(os.Getenv("PORT")); p != "" {
		return p
	}
	return "8080"
}

func newApp() *fiber.App {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(recover.New())

	g := app.Group(basePath())
	g.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
	g.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"service": "fiber-template", "base_path": basePath()})
	})
	return app
}

func main() {
	addr := ":" + port()
	log.Printf("fiber-template listening on %s (base_path=%q)", addr, basePath())
	if err := newApp().Listen(addr); err != nil {
		log.Fatal(err)
	}
}
