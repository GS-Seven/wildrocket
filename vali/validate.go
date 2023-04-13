package vali

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh2 "github.com/go-playground/validator/v10/translations/zh"
	"log"
)

var msg string

func BindValidate(params interface{}) string {
	translator, validate := validate()
	err := validate.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			msg = err.Translate(translator)
		}
		return msg
	}
	return ""
}

func validate() (ut.Translator, *validator.Validate) {
	v := validator.New()
	e := en.New()
	u := ut.New(e, e, zh.New())
	t, _ := u.GetTranslator("zh")
	err := zh2.RegisterDefaultTranslations(v, t)
	if err != nil {
		log.Printf("err: %v", err)
	}
	return t, v
}
