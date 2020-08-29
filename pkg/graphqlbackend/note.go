package graphqlbackend

import (
	"github.com/graphql-go/graphql"
)

// handles all graphql and services logic

type note struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NoteType() *graphql.Object {
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
