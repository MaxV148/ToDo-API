package api

import (
	db "CheckToDoAPI/db/sqlc"
	"CheckToDoAPI/middlewares"
	"CheckToDoAPI/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	queries        *db.Queries
	router         *gin.Engine
	tokenGenerator *utils.TokenGenerator
}

func NewServer(store *db.Queries, config utils.Config) *Server {
	tokenGen := utils.NewTokenGenerator(config)
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowOrigins = []string{"http://localhost:8000"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	server := &Server{queries: store, tokenGenerator: tokenGen}
	router := gin.Default()
	// API-Routes
	// User
	router.Use(CORSMiddleware())
	router.POST("/register", server.registerUser)
	router.POST("/login", server.loginUser)
	// Task
	protected := router.Group("/", middlewares.JwtAuthMiddleware())
	protected.POST("/todo", server.createToDo)
	protected.GET("/todos", server.getAllToDosForCurrentUser)
	protected.PUT("/todo", server.updateToDoFromCurrentUser)
	protected.POST("/done", server.makeToDoDoneFromCurrentUser)
	protected.DELETE("/todo", server.deleteToDoForCurrentUser)
	//categories
	protected.GET("/categories", server.getAllCategoriesForCurrentUser)
	protected.POST("/category", server.createCategory)
	//users
	protected.GET("/users", server.listAllUsers)
	//permissions
	protected.POST("/permissions", server.grandUserToDo)

	server.router = router
	return server
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
