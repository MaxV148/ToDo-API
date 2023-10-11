package api

import (
	db "CheckToDoAPI/db/sqlc"
	"CheckToDoAPI/middlewares"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type createToDoRequest struct {
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Category int64  `json:"category" binding:"required"`
}

type updateToDoRequest struct {
	ID      int64  `json:"todoId" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type makeToDoDoneRequest struct {
	ID int64 `json:"todoId" binding:"required"`
}

func (server *Server) createToDo(ctx *gin.Context) {
	var req createToDoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
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
	sortingOrder := ctx.DefaultQuery("sorting_order", "TITLE_ASC")
	todos, err := server.queries.ListToDoForUser(ctx, db.ListToDoForUserParams{UserID: currentUser, SortingOrder: sortingOrder})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, todos)

}

func (server *Server) updateToDoFromCurrentUser(ctx *gin.Context) {
	var req updateToDoRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}
	updated, err := server.queries.UpdateToDo(ctx, db.UpdateToDoParams{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, updated)
}

func (server *Server) makeToDoDoneFromCurrentUser(ctx *gin.Context) {
	var req makeToDoDoneRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}
	todo, err := server.queries.ToggleToDoDone(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, todo)
}

func (server *Server) deleteToDoForCurrentUser(ctx *gin.Context) {
	todoIdStr := ctx.Query("id")
	currentUser := ctx.GetInt64(middlewares.UserIDFromToken)

	if len(todoIdStr) == 0 {
		ctx.JSON(http.StatusBadRequest, "No id supplied")
		ctx.Abort()
		return
	}
	toDoId, err := strconv.Atoi(todoIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}

	_, err = server.queries.DeleteToDo(ctx, db.DeleteToDoParams{ID: int64(toDoId), CreatedBy: currentUser})
	log.Println("ERROR: ", err)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusBadRequest, "Not allowed to delete ist ToDo")
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.Status(http.StatusOK)
}
