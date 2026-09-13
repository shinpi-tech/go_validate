package custom

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Slug() validator.Func {
	return func(fl validator.FieldLevel) bool {
		s := fl.Field().String()
		matched, err := regexp.MatchString(`^[a-z0-9-_.]+$`, s)
		if err != nil {
			return false
		}
		return matched
	}
}

func SlugTranslation(ut ut.Translator) error {
	return ut.Add("slug", "{0} должен быть ярлыком", true)
}

func SlugValidation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("slug", fe.Field())
	return t
}
