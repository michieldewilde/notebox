package graphqlbackend

import (
	"github.com/graphql-go/graphql"
)

type note struct {
	ID      string `json:"id"`
	title   string `json:"title"`
	content string `json:"content"`
}

func UserType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Note",
			Fields: graphql.Fields{
				"id": &graphql.Field{
					Type: graphql.ID,
				},
				"title": &graphql.Field{
					Type: graphql.String,
				},
				"content": &graphql.Field{
					Type: graphql.String,
				},
			},
		},
	)
}
