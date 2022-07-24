package router

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

// Path of route
const (
	QueryPath      = "/query"
	PlaygroundPath = "/playground"
)

func New(srv *handler.Server) *gin.Engine {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(CORSMiddleware())

	router.GET("/playground", playgroundHandler())
	router.POST("/query", graphqlHandler(srv))

	return router
}

// Defining the Graphql handler
func graphqlHandler(srv *handler.Server) gin.HandlerFunc {
	// // NewExecutableSchema and Config are in the generated.go file
	// // Resolver is in the resolver.go file
	// h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
	// 	Resolvers: &resolver.Resolver{},
	// }))

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining CORS middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
