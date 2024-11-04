package validate_config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PhoneValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	pattern := `^\d{3}-\d{3}-\d{4}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(phone)
}
