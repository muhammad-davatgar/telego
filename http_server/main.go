package main

import (
	"context"
	"log"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/mmddvg/telego/http_server/utils"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {

	ctx := context.Background()

	neo, err := neo4j.NewDriverWithContext(os.Getenv("DBHost"), neo4j.BasicAuth(os.Getenv("DBUser"), os.Getenv("DBPass"), ""))

	if err != nil {
		log.Fatal("can't create driver : ", err)
	}

	if err = neo.VerifyConnectivity(ctx); err != nil {
		log.Fatal("couldn't connect : ", err)
	}

	e := echo.New()

	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	setup_routes(*e, &neo)

	e.Logger.Fatal(e.Start(":1323"))
}
