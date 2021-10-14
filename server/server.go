package server

import (
	"books-api/server/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   os.Getenv("APP_PORT"),
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	fmt.Println("Server is running 🚀")
	log.Fatal(router.Run(":" + s.port))
}
