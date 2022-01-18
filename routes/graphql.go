package routes

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/culy247/go-gin-template/graphql"
	"github.com/culy247/go-gin-template/graphql/generated"
	"github.com/gin-gonic/gin"
)

func Graphql(router *gin.Engine) *gin.Engine {

	router.POST("/graphql", graphqlHandler())
	router.GET("/graphiql", playgroundHandler())

	return router
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
