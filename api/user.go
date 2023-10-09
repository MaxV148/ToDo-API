package api

import (
	db "CheckToDoAPI/db/sqlc"
	"CheckToDoAPI/utils"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type userRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Username string `json:"username"`
	UserID   int64  `json:"user-id"`
	BaseResponse
}

type userResponse struct {
	ID       int64
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
	token, expiresAt, err := server.tokenGenerator.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, loginResponse{
		Username: user.Username,
		UserID:   user.ID,
		BaseResponse: BaseResponse{
			ExpiresAt: expiresAt,
			Token:     token,
		},
	})
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req userRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	user, err := server.queries.GetUserByName(ctx, req.Username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Println("user name not in DB: ", req.Username)
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		log.Printf("password for user: %s not correct.\n", req.Username)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, expiredAt, err := server.tokenGenerator.GenerateToken(user.ID)
	//utils.SetTokenHeader(ctx, accessToken)
	ctx.JSON(http.StatusOK, loginResponse{
		Username: user.Username,
		UserID:   user.ID,
		BaseResponse: BaseResponse{
			ExpiresAt: expiredAt,
			Token:     accessToken,
		},
	})
}

func (server *Server) listAllUsers(ctx *gin.Context) {
	users, err := server.queries.ListAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	userResponses := make([]userResponse, len(users))
	for i, user := range users {
		userResponses[i] = userResponse{
			ID:       user.ID,
			Username: user.Username,
		}
	}
	ctx.JSON(http.StatusOK, userResponses)
}
