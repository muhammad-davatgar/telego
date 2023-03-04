package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mmddvg/telego/http_server/utils"
)

func main() {
	e := echo.New()

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	setup_routes(*e)

	e.Logger.Fatal(e.Start(":1323"))
}
