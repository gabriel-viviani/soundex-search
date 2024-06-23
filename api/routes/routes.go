package routes

import (
	"github.com/Sigma-Ratings/sigma-code-challenges/api/controller"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// AttachRoutes binds controller functions to server routes
func AttachRoutes(e *echo.Echo, database *gorm.DB) {
	c := controller.Controller{DB: database}

	e.GET("/status", c.Status)
	e.GET("/search", c.Search)
}
