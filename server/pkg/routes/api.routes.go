package routes

import (
	"url_shortener/pkg/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoute(server *fiber.App){
	routes:=server.Group("")

	routes.Post("api/url",controller.HandleCreateShortURL)
	routes.All("redirect/:id",controller.HandleRedirect)
}