package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tiagoinacio/go-gql-server/internal/handlers"
	"github.com/tiagoinacio/go-gql-server/pkg/utils"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = utils.GetEnvVariable("GQL_SERVER_HOST")
	port = utils.GetEnvVariable("GQL_SERVER_PORT")
	gqlPath = utils.GetEnvVariable("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.GetEnvVariable("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.GetEnvVariableAsBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED")
}

func Run() {
	domain := host + ":" + port
	endpoint := "http://" + domain

	router := gin.Default()

	// Handlers
	// Keep-Alive/Ping handler
	router.GET("/ping", handlers.Ping())

	// GraphQL handlers
	if isPgEnabled {
		router.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Println("GraphQL Playground @ " + endpoint + gqlPath)
	}
	router.POST(gqlPath, handlers.GraphqlHandler())

	log.Println("GraphQL @ " + endpoint + gqlPath)

	// Run the server
	// Inform the user where the server is listening to
	log.Println("Running @ http://" + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatalln(router.Run(domain))
}
