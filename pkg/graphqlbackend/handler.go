package graphqlbackend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/friendsofgo/graphiql"
	"github.com/graphql-go/graphql"
)

type postData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}

type queryParams struct {
	Context   context.Context
	Schema    graphql.Schema
	Query     string
	Operation string
	Variables map[string]interface{}
}

func Handle(schema *graphql.Schema) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p postData
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			w.WriteHeader(400)
			return
		}

		params := &queryParams{
			Context:   r.Context(),
			Schema:    *schema,
			Query:     p.Query,
			Variables: p.Variables,
			Operation: p.Operation,
		}

		result := executeQuery(params)
		json.NewEncoder(w).Encode(result)
	}
}

func executeQuery(p *queryParams) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Context:        p.Context,
		Schema:         p.Schema,
		RequestString:  p.Query,
		VariableValues: p.Variables,
		OperationName:  p.Operation,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v\n", result.Errors)
	}
	return result
}

func HandleGraphiql(r string) http.Handler {
	graphiqlHandler, err := graphiql.NewGraphiqlHandler(r)
	if err != nil {
		fmt.Printf("Error serving graphiql: %v\n", err)
	}

	return graphiqlHandler
}
