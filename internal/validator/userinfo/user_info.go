package userinfo
import (
	"github.com/go-playground/validator/v10"
)

func NameValid(fl validator.FieldLevel) bool {
	name, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	} else {
		return len(name) >= 2 && len(name) <= 20
	}
}