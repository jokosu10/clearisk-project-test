package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HomeAPI(c echo.Context) error {
	return c.String(http.StatusOK, "Product API. API is Active")
}
