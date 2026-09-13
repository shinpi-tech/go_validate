package custom

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Time() validator.Func {
	return func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		// формат ЧЧ:ММ (00:00 - 23:59)
		matched, err := regexp.MatchString(`^([01][0-9]|2[0-3]):[0-5][0-9]$`, s)
		if err != nil {
			return false
		}
		return matched
	}
}

func TimeTranslation(ut ut.Translator) error {
	return ut.Add("time", "{0} должен иметь формат 00:00", true)
}

func TimeValidation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("time", fe.Field())
	return t
}
