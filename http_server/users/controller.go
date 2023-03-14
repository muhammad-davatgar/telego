package users

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func GetUser(neo *neo4j.DriverWithContext) echo.HandlerFunc {

	return func(c echo.Context) error {
		username := c.Param("username")

		user, err := GetUserQuery(neo, username)

		if err != nil {
			log.Println(err)
			return echo.ErrNotFound
		}

		return c.JSON(http.StatusOK, user.Props)
	}
}
func SignUp(neo *neo4j.DriverWithContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		validated, err_map := ValidateSignUp(&c)
		if len(err_map.Errors) != 0 {
			return echo.NewHTTPError(http.StatusBadRequest, err_map.Errors)
		}

		user, err := SignUpUserQuery(neo, validated)
		if err != nil {
			log.Println(err)
		}

		return c.JSON(http.StatusOK, user.Props)
	}
}
