package common

import (
	"unicode"

	"github.com/aliblue2/khodro45/configs"
)

// var (
// 	lowerCharSet   = "abcdedfghijklmnopqrstuvwxyz"
// 	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	specialCharSet = "!@#$%&*"
// 	numberSet      = "0123456789"
// 	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
// )

func CheckPassword(password string, cfg *configs.Config) bool {

	if len(password) < cfg.Password.MinLength {
		return false
	}
	if cfg.Password.IncludeChars && !HasLetter(password) {
		return false
	}
	if cfg.Password.IncludeDigits && !HasDigits(password) {
		return false
	}
	if cfg.Password.IncludeLowercase && !HasLower(password) {
		return false
	}
	if cfg.Password.IncludeUppercase && !HasUpper(password) {
		return false
	}
	return true

}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasDigits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true

		}
	}
	return false
}

func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}

	return false
}

func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}

	return false
}
