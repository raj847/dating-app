package validate

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.RegisterValidation("passwd", IsPasswordValid)
	validate.RegisterValidation("gendr", ValidateGender)
}

func Validate(v interface{}) error {
	return validate.Struct(v)
}

func IsPasswordValid(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	var hasLowercase, hasUppercase, hasNumber, hasSymbol bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSymbol = true
		}
	}

	return hasUppercase && hasLowercase && hasNumber && hasSymbol
}

func ValidateGender(fl validator.FieldLevel) bool {
	gender := fl.Field().String()
	return gender == "L" || gender == "P"
}
