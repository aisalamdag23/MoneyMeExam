package customvalidators

import (
	"gopkg.in/go-playground/validator.v9"
)

var DateOfBirthValidator validator.Func = func(fl validator.FieldLevel) bool {
	// val := fl.Field().String()
	// timeVal, err := time.Parse("2006-01-02", val)
	// if err != nil {
	// 	return false
	// }
	// now := time.Now()
	// return !now.Before(timeVal)
	return true
}
