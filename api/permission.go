package api

import (
	db "CheckToDoAPI/db/sqlc"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				ctx.Abort()
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.Status(http.StatusOK)
}
