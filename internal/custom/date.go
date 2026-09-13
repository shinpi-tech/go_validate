package custom

import (
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Date() validator.Func {
	return func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		_, err := time.Parse("2006-01-02", s)
		return err == nil
	}
}

func DateTranslation(ut ut.Translator) error {
	return ut.Add("date", "{0} должен быть в формате ГГГГ-ММ-ДД", true)
}

func DateValidation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("date", fe.Field())
	return t
}
