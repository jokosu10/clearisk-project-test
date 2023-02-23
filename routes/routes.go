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

	e.GET("/test/organizations", controllers.TestOrganizatiosAPI)
	e.GET("/organizations", controllers.GetOrganizations)
	e.POST("/organization", controllers.CreateOrganization)
	e.DELETE("/organization/:id", controllers.DeleteOrganization)

	e.GET("/test/payments", controllers.TestPaymentsAPI)
	// e.GET("/organizations", controllers.GetOrganizations)
	// e.POST("/organization", controllers.CreateOrganization)
	// e.DELETE("/organization/:id", controllers.DeleteOrganization)

	return e
}
