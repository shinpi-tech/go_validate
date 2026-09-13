package validate

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/shinpi-tech/go_validate/internal/custom"

	rus "github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ru_translations "github.com/go-playground/validator/v10/translations/ru"
)

// Valid — обёртка над validator с русской локализацией ошибок.
type Valid struct {
	validate *validator.Validate
	trans    ut.Translator
}

func (h *Valid) Validate(s any) error {
	err := h.validate.Struct(s)
	if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)

		var result string

		for i, e := range errs {
			result += e.Translate(h.trans)

			if i < len(errs)-1 {
				result += "\n"
			}
		}

		return fmt.Errorf(`%s`, result)
	}

	return nil
}

// Custom описывает дополнительное правило валидации, специфичное для сервиса.
type Custom struct {
	Tag                 string
	Func                validator.Func
	RegisterTranslation func(ut ut.Translator) error
	Translation         validator.TranslationFunc
}

// NewValid создаёт валидатор с общими русскими переводами и встроенными правилами.
// Через extra можно добавить сервис-специфичные правила.
func NewValid(extra ...Custom) *Valid {
	ru := rus.New()
	uni := ut.New(ru, ru)
	trans, _ := uni.GetTranslator("ru")

	validateInit := validator.New(validator.WithRequiredStructEnabled())

	validateInit.RegisterTagNameFunc(func(fld reflect.StructField) string { return fld.Tag.Get("name") })

	validateInit.RegisterValidation("only_num", custom.OnlyNum())
	validateInit.RegisterValidation("slug", custom.Slug())
	validateInit.RegisterValidation("snake", custom.Snake())
	validateInit.RegisterValidation("phone", custom.Phone())
	validateInit.RegisterValidation("date", custom.Date())
	validateInit.RegisterValidation("time", custom.Time())

	validateInit.RegisterTranslation("only_num", trans, custom.OnlyNumTranslation, custom.OnlyNumValidation)
	validateInit.RegisterTranslation("slug", trans, custom.SlugTranslation, custom.SlugValidation)
	validateInit.RegisterTranslation("snake", trans, custom.SnakeTranslation, custom.SnakeValidation)
	validateInit.RegisterTranslation("phone", trans, custom.PhoneTranslation, custom.PhoneValidation)
	validateInit.RegisterTranslation("date", trans, custom.DateTranslation, custom.DateValidation)
	validateInit.RegisterTranslation("time", trans, custom.TimeTranslation, custom.TimeValidation)

	for _, c := range extra {
		if c.Tag == "" || c.Func == nil {
			continue
		}
		validateInit.RegisterValidation(c.Tag, c.Func)
		if c.RegisterTranslation != nil && c.Translation != nil {
			_ = validateInit.RegisterTranslation(c.Tag, trans, c.RegisterTranslation, c.Translation)
		}
	}

	err := ru_translations.RegisterDefaultTranslations(validateInit, trans)
	if err != nil {
		return nil
	}

	return &Valid{
		validate: validateInit,
		trans:    trans,
	}
}
