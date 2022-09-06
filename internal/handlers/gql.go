package handlers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/tiagoinacio/go-gql-server/internal/gql"
	"gitub.com/tiagoinacio/go-gql-server/internal/gql/resolvers"
)

func GraphqlHandler() gin.HandlerFunc {
	c := gql.Config{
		Resolvers: &resolvers.Resolver{},
	}

	h := handler.GraphQL(gql.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := handler.Playground("Go GraphQL Server", path)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
