package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mmddvg/telego/http_server/users"
)

func setup_routes(app echo.Echo) {
	users_group := app.Group("/users")
	users_group.GET("/:id", users.GetUser)
	users_group.POST("", users.SignUp)
}
