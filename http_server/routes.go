package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mmddvg/telego/http_server/users"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func setup_routes(app echo.Echo, neo *neo4j.DriverWithContext) {
	users_group := app.Group("/users")
	users_group.GET("/:username", users.GetUser(neo))
	users_group.POST("", users.SignUp(neo))
}
