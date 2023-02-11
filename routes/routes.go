package routes

import (
	"project-alta-store/controllers"

	"github.com/labstack/echo"
)

func Start() *echo.Echo {
	e := echo.New()

	// route swagger ui
	e.Static("/swagger-ui.css", "dist/swagger-ui.css")
	e.Static("/swagger-ui-bundle.js", "dist/swagger-ui-bundle.js")
	e.Static("/swagger-ui-standalone-preset.js", "dist/swagger-ui-standalone-preset.js")
	e.Static("/swagger.yaml", "swagger.yaml")
	e.Static("/swaggerui", "dist/index.html")

	e.GET("/test", controllers.HomeAPI)
	e.GET("/organizations", controllers.OrganizationAPI)
	e.POST("/organization", controllers.CreateOrganization)
	return e
}
