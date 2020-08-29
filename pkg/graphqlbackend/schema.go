package graphqlbackend

import (
	"github.com/graphql-go/graphql"
)

func NewSchema() (*graphql.Schema, error) {
	queryType := graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"note": &graphql.Field{
					Type: UserType(),
					Args: graphql.FieldConfigArgument{
						"id": &graphql.ArgumentConfig{
							Type: graphql.String,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return &note{
							ID:      "test",
							title:   "title",
							content: "content",
						}, nil
					},
				},
			},
		},
	)

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
