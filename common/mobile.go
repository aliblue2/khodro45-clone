package common

import (
	"log"
	"regexp"
)

const iranMobileNumberPattern = `(09(0[1-5]|1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7})`

func ValidatePhoneNumber(number string) bool {
	result, err := regexp.MatchString(iranMobileNumberPattern, number)

	if err != nil {
		log.Fatalf("invalid phone number")
		return false
	}
	return result

}
