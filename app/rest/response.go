package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ilya-sokolov/crypto_kiddies-server/app/rest/check"
	"net/http"
)

type ErrorMessage struct {
	Message string                                 `json:"message"`
	Fields  validator.ValidationErrorsTranslations `json:"fields,omitempty"`
	Details string                                 `json:"details,omitempty"`
}

func ResponseSuccess(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(code, data)
}

func ResponseError(ctx *gin.Context, code int, err ErrorMessage) {
	ctx.JSON(code, err)
}

func ResponseValidationError(ctx *gin.Context, err validator.ValidationErrors) {
	fields := check.Translate(ctx, err)
	ResponseError(ctx, http.StatusBadRequest, ErrorMessage{Message: "Validation error", Fields: fields})
}
