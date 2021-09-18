package utils

import (
	"github.com/go-playground/validator/v10"
	"unicode"
	"unicode/utf8"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Init() {
	cv.Validator.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		var (
			hasNumber      = false
			hasSpecialChar = false
			hasLetter      = false
			hasSuitableLen = false
		)

		password := fl.Field().String()

		if utf8.RuneCountInString(password) <= 30 || utf8.RuneCountInString(password) >= 6 {
			hasSuitableLen = true
		}

		for _, c := range password {
			switch {
			case unicode.IsNumber(c):
				hasNumber = true
			case unicode.IsPunct(c) || unicode.IsSymbol(c):
				hasSpecialChar = true
			case unicode.IsLetter(c) || c == ' ':
				hasLetter = true
			default:
				return false
			}
		}

		return hasNumber && hasSpecialChar && hasLetter && hasSuitableLen
	})
}


func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}