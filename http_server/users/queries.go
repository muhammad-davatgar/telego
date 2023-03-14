package users

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

func GetUserQuery(neo *neo4j.DriverWithContext, username string) (ValidatedUser, error) {
	ctx := context.Background()
	session := (*neo).NewSession(ctx, neo4j.SessionConfig{})

	data, err := neo4j.ExecuteRead(
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
		return ValidatedUser{}, fmt.Errorf("query error : %w", err)
	}

	user, err := ValidatedUserFromMap(data.Props)
	if err != nil {
		return ValidatedUser{}, fmt.Errorf("mapping error : %w", err)
	}

	return user, err
}

func SignUpUserQuery(neo *neo4j.DriverWithContext, user_entry ValidatedUser) (dbtype.Node, error) {
	ctx := context.Background()
	session := (*neo).NewSession(ctx, neo4j.SessionConfig{})

	user, err := neo4j.ExecuteWrite(
		ctx,
		session,
		func(tx neo4j.ManagedTransaction) (neo4j.Node, error) {
			encrypted_pass, err := EncryptPassword(user_entry.Username, user_entry.Password)

			if err != nil {
				return neo4j.Node{}, fmt.Errorf("encrypting : %w", err)
			}

			result, err := tx.Run(ctx,
				`MERGE (u:User {username : $username})
					set u.password = $password
					return u`,
				map[string]any{"username": user_entry.Username, "password": encrypted_pass},
			)
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

	return user, err
}
