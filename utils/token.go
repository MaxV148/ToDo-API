package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"strconv"
	"time"
)

const (
	USERID = "userID"
)

type TokenGenerator struct {
	config Config
}

func SetTokenHeader(ctx *gin.Context, token string) {
	ctx.Writer.Header().Set("Authorization", token)
}

func NewTokenGenerator(config Config) *TokenGenerator {
	return &TokenGenerator{config: config}
}

func (t *TokenGenerator) GenerateToken(userId int64) (string, error) {

	ttl, _ := strconv.Atoi(t.config.TokenLifespan)

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims[USERID] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(ttl)).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.config.APISecret))
}

func TokenValid(token string) (int64, error) {
	config, err := LoadConfig("../")
	if err != nil {
		log.Fatalln("cannot load config: ", err)
	}

	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.APISecret), nil
	})

	if err != nil {
		return -1, err
	}
	claims, ok := tok.Claims.(jwt.MapClaims)
	if ok && tok.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims[USERID]), 10, 32)
		if err != nil {
			return -1, err
		}
		return int64(uid), nil
	}

	return -1, nil
}
