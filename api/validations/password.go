package validations

import (
	"log"

	"github.com/aliblue2/khodro45/common"
	"github.com/aliblue2/khodro45/configs"
	"github.com/go-playground/validator/v10"
)

func CheckPassword(fld validator.FieldLevel) bool {
	val, ok := fld.Field().Interface().(string)

	if !ok {
		log.Println("invalida password type")
		return false
	}

	cfg := configs.GetConfig()
	result := common.CheckPassword(val, cfg)
	return result

}
