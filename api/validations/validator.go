package validations

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

//(09(0[1-5]|1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7})

func IranPhoneNumChecker(fld validator.FieldLevel) bool {
	val, ok := fld.Field().Interface().(string)

	if !ok {
		log.Fatalln("invalid phone number")
		return false
	}

	res, err := regexp.MatchString(`(09(0[1-5]|1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7})`, val)

	if err != nil {
		return false
	}
	return res
}
