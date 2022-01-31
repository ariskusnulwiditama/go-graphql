package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	// "mygraphql/graphql/query"
)

func NewHandler() (*handler.Handler, error) {
	shcema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: newQuery(),
		},
	)
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:   &shcema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}
