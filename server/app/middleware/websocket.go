package middleware

import (
	"log"
	"github.com/gofiber/fiber/v2"
)

func WebsocketCheckOrigin(c *fiber.Ctx) error {
	log.Println("websocket check origin")
	return c.Next()
	//if c.Get("host") == "localhost:9097" {
	//	c.Locals("Host", "Localhost:9097")
	//	return c.Next()
	//}
	//return c.Status(403).SendString("Request origin not allowed middleware")
}
