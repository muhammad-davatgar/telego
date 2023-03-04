package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func SignUp(c echo.Context) error {
	validated, err := ValidateSignUp(&c)
	if len(err.Errors) != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, err.Errors)
	}
	return c.JSON(http.StatusOK, validated)
}
