package custom

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func OnlyNum() validator.Func {
	return func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		for _, c := range s {
			if c < '0' || c > '9' {
				return false
			}
		}
		return true
	}
}

func OnlyNumTranslation(ut ut.Translator) error {
	return ut.Add("only_num", "{0} должен содержать только цифры", true)
}

func OnlyNumValidation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("only_num", fe.Field())
	return t
}
