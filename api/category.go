package api

import (
	"CheckToDoAPI/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) getAllCategoriesForCurrentUser(ctx *gin.Context) {
	currentUser := ctx.GetInt64(middlewares.UserIDFromToken)
	cats, err := server.queries.ListCategoriesForUser(ctx, currentUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, cats)

}
