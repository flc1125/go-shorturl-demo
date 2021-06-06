package routes

import (
	"shorturl/app/controllers"
	"shorturl/app/controllers/shorturl"

	"github.com/labstack/echo/v4"
)

func Web(e *echo.Echo) {
	e.GET("/", controllers.Hello)
	e.POST("/create", shorturl.Create)
	e.GET("/:url", shorturl.Redirect)
}
