package main

import (
	"github.com/andersoncarubelli/go-fiber-rest-api/pkg/configs"
	"github.com/andersoncarubelli/go-fiber-rest-api/pkg/middleware"
	"github.com/andersoncarubelli/go-fiber-rest-api/pkg/routes"
	"github.com/andersoncarubelli/go-fiber-rest-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@email.com
// @license.name Apchae 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apkiKey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.NotFoundRoute(app)

	utils.StartServerWithGracefulShutdown(app)
}
