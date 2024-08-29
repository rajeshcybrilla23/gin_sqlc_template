package controller

import (
	service "gin-template/api_v2/service"
	db "gin-template/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountController interface {
	createAccount(ctx *gin.Context)
}

type AccountControllerImpl struct {
	accService service.AccountServiceImpl
}

func AccountControllerImplInit(accService service.AccountServiceImpl) AccountControllerImpl {
	return AccountControllerImpl{
		accService: accService,
	}
}

func (c AccountControllerImpl) CreateAccount(ctx *gin.Context) {
	account, err := c.accService.CreateAccount(ctx)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, account)
}
