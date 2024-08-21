package api

import (
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency" binding:"required,oneof=USD INR"`
}

func (server *Server) createAccount(ctx *gin.Context) {

}
