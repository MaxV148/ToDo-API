package middlewares

import (
	"CheckToDoAPI/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	AuthHeaderKey   = "Authorization"
	AuthTypeBearer  = "bearer"
	UserIDFromToken = "user_id_payload"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthHeaderKey)

		if len(authHeader) == 0 {
			err := errors.New("authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		fields := strings.Split(authHeader, " ")
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		authType := strings.ToLower(fields[0])
		if authType != AuthTypeBearer {
			err := errors.New("unsupported authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		accessToken := fields[1]
		uid, err := utils.TokenValid(accessToken)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized, Bad Token")
			c.Abort()
			return
		}
		// set Userid in context
		c.Set(UserIDFromToken, uid)

		c.Next()
	}
}
