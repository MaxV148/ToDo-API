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

type updateToDoRequest struct {
	ID int64 `json:"todoId" binding:"required"`
	createToDoRequest
}

type makeToDoDoneRequest struct {
	ID int64 `json:"todoId" binding:"required"`
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

func (server *Server) updateToDoFromCurrentUser(ctx *gin.Context) {
	var req updateToDoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	updated, err := server.queries.UpdateToDo(ctx, db.UpdateToDoParams{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, updated)
}

func (server *Server) makeToDoDoneFromCurrentUser(ctx *gin.Context) {
	var req makeToDoDoneRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	todo, err := server.queries.ToggleToDoDone(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, todo)
}
