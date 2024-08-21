package api

import (
	db "gin-template/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	BaseController
}

func (accountController *AccountController) createAccount(ctx *gin.Context) {

	account, err := accountController.services.AccountService.Create(ctx)
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
