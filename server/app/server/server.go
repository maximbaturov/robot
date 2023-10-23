package server

import (
	"log"
	"robot/app/config"
	"robot/app/controllers"
	"robot/app/middleware"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	config := config.ParseConfig()

	app := fiber.New()

	app.Use(recover.New())

	app.Use("/websocket", middleware.WebsocketCheckOrigin)
	app.Get("/websocket", websocket.New(controllers.Websocket))

	app.Static("/", "./web/dist/")
	app.Static("/*", "./web/dist/index.html")

	log.Fatal(app.Listen(config.HttpPort))

}
