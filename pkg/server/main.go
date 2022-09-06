package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tiagoinacio/go-gql-server/internal/handlers"
	"github.com/tiagoinacio/go-gql-server/pkg/utils"
)

var host, port string

func init() {
	host = utils.GetEnvVariable("GQL_SERVER_HOST")
	port = utils.GetEnvVariable("GQL_SERVER_PORT")
}

func Run() {
	router := gin.Default()
	// Keep-Alive/Ping handler
	router.GET("/ping", handlers.Ping())
	// Inform the user where the server is listening to
	log.Println("Running @ http://" + host + ":" + port)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatalln(router.Run(host + ":" + port))
}
