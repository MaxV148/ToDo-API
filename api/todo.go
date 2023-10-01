package api

import (
	db "CheckToDoAPI/db/sqlc"
	"CheckToDoAPI/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createToDoRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Category int64  `json:"category" binding:"required"`
}

func (server *Server) createToDo(ctx *gin.Context) {
	var req createToDoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	currentUser := ctx.GetInt64(middlewares.UserIDFromToken)

	todo, err := server.queries.CreateToDo(ctx, db.CreateToDoParams{
		CreatedBy: currentUser,
		Title:     req.Title,
		Content:   req.Content,
		Category:  req.Category,
	})
	if err != nil {
		// TODO: handle database error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, todo)
}

func (server *Server) getAllToDosForCurrentUser(ctx *gin.Context) {
	currentUser := ctx.GetInt64(middlewares.UserIDFromToken)
	todos, err := server.queries.ListToDoForUser(ctx, currentUser)
	if err != nil {
		// TODO: handle database error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, todos)

}
