package api

import (
	db "CheckToDoAPI/db/sqlc"
	"CheckToDoAPI/utils"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type createUserResponse struct {
	Username string
}

func (server *Server) registerUser(ctx *gin.Context) {
	var req userRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	hashedPw, err := utils.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	arg := db.CreateUserParams{Username: req.Username, Password: hashedPw}
	user, err := server.queries.CreateUser(ctx, arg)
	if err != nil {
		// TODO: handle database error
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	token, err := server.tokenGenerator.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	resp := createUserResponse{Username: user.Username}
	utils.SetTokenHeader(ctx, token)
	ctx.JSON(http.StatusOK, resp)
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req userRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	user, err := server.queries.GetUserByName(ctx, req.Username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenGenerator.GenerateToken(user.ID)
	utils.SetTokenHeader(ctx, accessToken)
	ctx.Status(http.StatusOK)
}
