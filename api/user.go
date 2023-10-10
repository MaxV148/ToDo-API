package api

import (
	db "CheckToDoAPI/db/sqlc"
	"CheckToDoAPI/utils"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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
		ctx.Abort()
		return
	}
	arg := db.CreateUserParams{Username: req.Username, Password: hashedPw}
	user, err := server.queries.CreateUser(ctx, arg)
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
	token, expiresAt, err := server.tokenGenerator.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
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
		ctx.Abort()
		return
	}
	user, err := server.queries.GetUserByName(ctx, req.Username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			log.Println("user name not in DB: ", req.Username)
			ctx.JSON(http.StatusUnauthorized, errorResponse(err))
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		ctx.Abort()
		return
	}
	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		log.Printf("password for user: %s not correct.\n", req.Username)
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		ctx.Abort()
		return
	}

	accessToken, expiredAt, err := server.tokenGenerator.GenerateToken(user.ID)
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
