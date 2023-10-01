package api

import (
	db "CheckToDoAPI/db/sqlc"
	"CheckToDoAPI/middlewares"
	"CheckToDoAPI/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	queries        *db.Queries
	router         *gin.Engine
	tokenGenerator *utils.TokenGenerator
}

func NewServer(store *db.Queries, config utils.Config) *Server {
	tokenGen := utils.NewTokenGenerator(config)
	server := &Server{queries: store, tokenGenerator: tokenGen}
	router := gin.Default()
	// API-Routes
	// User
	router.POST("/register", server.registerUser)
	router.POST("/login", server.loginUser)
	// Task
	protected := router.Group("/auth")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.POST("/todo", server.createToDo)
	protected.GET("/todos", server.getAllToDosForCurrentUser)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
