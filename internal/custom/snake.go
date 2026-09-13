package custom

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Snake() validator.Func {
	return func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		return regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(value)
	}
}

func SnakeTranslation(ut ut.Translator) error {
	return ut.Add("snake", "{0} может содержать только латинские буквы нижнего регистра, цифры и _", true)
}

func SnakeValidation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("snake", fe.Field())
	return t
}
