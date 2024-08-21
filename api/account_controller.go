package api

import (
	db "gin-template/db/sqlc"
	"net/http"
	service "gin-template/service"
	"github.com/gin-gonic/gin"
)



func createAccount(ctx *gin.Context) {

	account, err := service.
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
