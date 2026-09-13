package custom

import (
	"regexp"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func Phone() validator.Func {
	return func(fl validator.FieldLevel) bool {
		phone := fl.Field().String()
		// 10–15 цифр без знака "+"
		return regexp.MustCompile(`^[0-9]{10,15}$`).MatchString(phone)
	}
}

func PhoneTranslation(ut ut.Translator) error {
	return ut.Add("phone", "{0} должен содержать только цифры (10–15 знаков)", true)
}

func PhoneValidation(ut ut.Translator, fe validator.FieldError) string {
	t, _ := ut.T("phone", fe.Field())
	return t
}
