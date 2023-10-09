package api

import (
	db "CheckToDoAPI/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type grandToDo struct {
	UserIDToGrand int64 `json:"userToGrand" binding:"required"`
	ToDoForGrand  int64 `json:"ToDoForGrand" binding:"required"`
}

func (server *Server) grandUserToDo(ctx *gin.Context) {
	var req grandToDo
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}

	_, err := server.queries.GrantUserToToDo(ctx, db.GrantUserToToDoParams{UserID: req.UserIDToGrand, TodoID: req.ToDoForGrand})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.Status(http.StatusOK)
}
