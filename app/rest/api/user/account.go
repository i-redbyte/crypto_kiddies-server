package user

import (
	"github.com/gin-gonic/gin"
	"github.com/ilya-sokolov/crypto_kiddies-server/app/rest"
	"github.com/ilya-sokolov/crypto_kiddies-server/storage"
	"net/http"
)

func Account(acc *gin.RouterGroup) {
	// TODO: Red_byte Add account-owned routing
}

func CreateAccount(ctx *gin.Context) {
	type request struct {
		NickName string `json:"nickName" validate:"required,gt=1,lt=100"`
		Password string `json:"password" validate:"required,gt=5,lt=100"`
		Email    string `json:"email" validate:"required,email,gt=5,lt=100"`
	}
	var r request
	if err := rest.BindAndValidate(ctx, &r); err != nil {
		return
	}
	account, err := storage.CreateAccount(r.NickName, r.Email, r.Password)
	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, rest.ErrorMessage{Message: err.Error()})
		return
	}
	rest.ResponseSuccess(ctx, http.StatusCreated, AccountResponse{Account: *account, Token: account.Token()})
}

func Authorization(ctx *gin.Context) {
	type request struct {
		NickName string `json:"nickName" validate:"required,gt=0"`
		Password string `json:"password" validate:"required,gt=5,lt=100"`
	}
	var r request
	if err := rest.BindAndValidate(ctx, &r); err != nil {
		return
	}
	account, err := storage.GetAccount(r.NickName, r.Password)
	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, rest.ErrorMessage{Message: err.Error()})
		return
	}
	rest.ResponseSuccess(ctx, http.StatusOK, AccountResponse{Account: *account, Token: account.Token()})
}

type AccountResponse struct {
	storage.Account
	Token string `json:"token"`
}
