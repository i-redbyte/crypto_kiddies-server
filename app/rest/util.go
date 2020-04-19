package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ilya-sokolov/crypto_kiddies-server/app/rest/check"
	"net/http"
)

func BindAndValidate(ctx *gin.Context, r interface{}) error {
	err := ctx.BindJSON(&r)
	if err != nil {
		ResponseError(ctx, http.StatusBadRequest, ErrorMessage{Message: err.Error()})
		return err
	}
	err = check.Struct(r)
	if err != nil {
		ResponseValidationError(ctx, err.(validator.ValidationErrors))
		return err
	}
	return nil
}
