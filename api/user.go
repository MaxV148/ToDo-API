package api

import (
	db "CheckToDoAPI/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type createUserResponse struct {
	Username string
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	arg := db.CreateUserParams{Username: req.Username, Password: req.Password}
	user, err := server.queries.CreateUser(ctx, arg)
	if err != nil {
		// TODO: handle database error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	resp := createUserResponse{Username: user.Username}
	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) loginUser(ctx *gin.Context) {

}
