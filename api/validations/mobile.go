package validations

import (
	"log"

	"github.com/aliblue2/khodro45/common"
	"github.com/go-playground/validator/v10"
)

func IranPhoneNumChecker(fld validator.FieldLevel) bool {
	val, ok := fld.Field().Interface().(string)

	if !ok {
		log.Fatalf("invalida number type")
		return false
	}

	res := common.ValidatePhoneNumber(val)

	return res

}
