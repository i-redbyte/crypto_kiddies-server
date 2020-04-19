package check

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ent "github.com/go-playground/validator/v10/translations/en"
)

var validate *validator.Validate
var uni *ut.UniversalTranslator

func init() {
	english := en.New()
	uni = ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	validate = validator.New()
	err := ent.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
}

func Var(field interface{}, tags string) error {
	return validate.Var(field, tags)
}

func Struct(field interface{}) error {
	return validate.Struct(field)
}

func Translate(ctx *gin.Context, err validator.ValidationErrors) validator.ValidationErrorsTranslations {
	u, _ := uni.GetTranslator(ctx.GetHeader("Accept-Language"))
	return err.Translate(u)
}
