package users

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

func GetUser(neo *neo4j.DriverWithContext) echo.HandlerFunc {

	return func(c echo.Context) error {
		username := c.Param("username")

		ctx := context.Background()
		session := (*neo).NewSession(ctx, neo4j.SessionConfig{})

		user, err := neo4j.ExecuteRead(
			ctx,
			session,
			func(tx neo4j.ManagedTransaction) (dbtype.Node, error) {
				result, err := tx.Run(ctx, "MATCH (u:User {username : $username}) return u", map[string]any{"username": username})
				if err != nil {
					return *new(neo4j.Node), err
				}

				return neo4j.SingleTWithContext(ctx, result,
					func(record *neo4j.Record) (neo4j.Node, error) {
						node, _, err := neo4j.GetRecordValue[neo4j.Node](record, "u")
						return node, err
					},
				)
			},
		)

		if err != nil {
			log.Println(err)
			return echo.ErrNotFound
		}

		return c.JSON(http.StatusOK, user.Props)
	}
}
func SignUp(neo *neo4j.DriverWithContext) echo.HandlerFunc {
	return func(c echo.Context) error {
		validated, err := ValidateSignUp(&c)
		if len(err.Errors) != 0 {
			return echo.NewHTTPError(http.StatusBadRequest, err.Errors)
		}
		return c.JSON(http.StatusOK, validated)
	}
}
