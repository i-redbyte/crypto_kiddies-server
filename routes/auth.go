package routes

import (
	"github.com/gin-gonic/gin"
	. "github.com/ilya-sokolov/crypto_kiddies-server/common"
	"github.com/ilya-sokolov/crypto_kiddies-server/errors"
	"github.com/ilya-sokolov/crypto_kiddies-server/model"
	"net/http"
	"os"
)

func CheckToken(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		ResponseError(ctx, http.StatusUnauthorized, ErrorMessage{Message: errors.InvalidToken.Error()})
		ctx.Abort()
		return
	}
	token, err := model.ParseToken(auth, []byte(os.Getenv("token_password")))
	if err != nil {
		ResponseError(ctx, http.StatusUnauthorized, ErrorMessage{Message: errors.InvalidToken.Error()})
		ctx.Abort()
		return
	}
	account, err := model.GetAccountById(token.Id)
	if err != nil {
		ResponseError(ctx, http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
		ctx.Abort()
		return
	}
	ctx.Set("account", account)
	ctx.Set("accountId", account.Id)
	ctx.Next()
}
