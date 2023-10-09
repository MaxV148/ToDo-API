package api

import (
	db "CheckToDoAPI/db/sqlc"
	"CheckToDoAPI/middlewares"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type createCategory struct {
	Name string `json:"name" binding:"required"`
}

func (server *Server) getAllCategoriesForCurrentUser(ctx *gin.Context) {
	currentUser := ctx.GetInt64(middlewares.UserIDFromToken)
	sortOrder := ctx.DefaultQuery("sort_order", "TITLE_ASC")
	log.Println("GET all categories for User: ", db.ListToDoForUserParams{UserID: currentUser, SortingOrder: sortOrder})
	cats, err := server.queries.ListCategoriesForUser(ctx, currentUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, cats)
}

func (server *Server) createCategory(ctx *gin.Context) {
	var req createCategory
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		ctx.Abort()
		return
	}
	currentUser := ctx.GetInt64(middlewares.UserIDFromToken)
	log.Printf("CREATE category for User: %d ", currentUser)
	cat, err := server.queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: req.Name,
		User: currentUser,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, cat)
}
