package common

import (
	"log"
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
	if len(password) < cfg.Passwrod.MinLength {
		log.Println("because of minlength")
		return false
	}

	if cfg.Passwrod.IncludeChars && !HasLetter(password) {
		log.Println("because of chars")
		return false
	}

	if cfg.Passwrod.IncludeDigits && !HasDigits(password) {
		log.Println("because of digits")
		return false
	}

	if cfg.Passwrod.IncludeLowercase && !HasLower(password) {
		log.Println("because of lower")
		return false
	}

	if cfg.Passwrod.IncludeUppercase && !HasUpper(password) {
		log.Println("because of upper")
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
