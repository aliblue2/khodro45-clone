package common

import (
	"regexp"

	"github.com/aliblue2/khodro45/configs"
	"github.com/aliblue2/khodro45/pkg/logging"
)

const iranMobileNumberPattern = `(09(0[1-5]|1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7})`

var logger = logging.NewLogger(configs.GetConfig())

func ValidatePhoneNumber(number string) bool {
	result, err := regexp.MatchString(iranMobileNumberPattern, number)

	if err != nil {
		logger.Fatal(logging.Validation, logging.MobileValidation, err.Error(), nil)
		return false
	}
	return result

}
