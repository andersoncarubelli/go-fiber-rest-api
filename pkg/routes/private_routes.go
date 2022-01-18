package routes

import (
	"github.com/andersoncarubelli/go-fiber-rest-api/app/controllers"
	"github.com/andersoncarubelli/go-fiber-rest-api/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtectd(), controllers.CreateBook)
	route.Put("/book", middleware.JWTProtectd(), controllers.UpdateBook)
	route.Delete("/book", middleware.JWTProtectd(), controllers.DeleteBook)
}
