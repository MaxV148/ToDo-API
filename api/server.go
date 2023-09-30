package api

import (
	db "CheckToDoAPI/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(store *db.Queries) *Server {
	server := &Server{queries: store}
	router := gin.Default()
	// API-Routes
	// User
	router.POST("/register", server.createUser)
	router.POST("/login", server.loginUser)
	// Task
	router.POST("/todo")
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
