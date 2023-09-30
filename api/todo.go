package api

import "github.com/gin-gonic/gin"

type createToDoRequest struct {
	Title    string
	Content  string
	Category int64
}

func (server *Server) createToDo(ctx *gin.Context) {

}
