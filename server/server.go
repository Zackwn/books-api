package server

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zackwn/books-api/server/routes"
)

type Server struct {
	port   string
	server *gin.Engine
}

func New() Server {
	return Server{
		port:   "8080",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	// log.Println("server is running on port", s.port)
	log.Fatal(router.Run(":" + s.port))
}
