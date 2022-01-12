package utils

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/aaronangxz/TIC2601/models"
)

//Validate if a is > b
func ValidateLimitMax(a uint32, b uint32) bool {
	return a > b
}

func ValidateMaxStringLength(s string) bool {
	return len(s) < int(models.MaxStringLength)
}

func ValidateMaxItemNameStringLength(s string) bool {
	return len(s) < int(models.MaxItemNameStringLength)
}

func ValidateMaxItemDescriptionStringLength(s string) bool {
	return len(s) < int(models.MaxItemDescriptionStringLength)
}

func ValidateUint(a *uint32) bool {
	return *a >= uint32(0)
}

func ValidateInt64(a *int64) bool {
	return fmt.Sprint(reflect.TypeOf(a)) == "*int64"
}

func ValidateString(a *string) bool {
	return fmt.Sprint(reflect.TypeOf(a)) == "*string"
}

func IsContainsSpecialChar(a string) bool {
	for _, char := range a {
		if unicode.IsSymbol(char) {
			return true
		}
	}
	for _, char := range a {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return true
		}
	}
	return false
}

func IsContainsSpace(a string) bool {
	for _, char := range a {
		if unicode.IsSpace(char) {
			return true
		}
	}
	return false
}

func IsContainsAtSign(a string) bool {
	match := "@"

	for _, char := range a {
		for _, key := range match {
			if char == key {
				return true
			}
		}
	}
	return false
}

func RemoveSpecialChar(a string) string {
	for _, char := range a {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) && !unicode.IsSpace(char) {
			a = strings.ReplaceAll(a, string(char), "")
		}
	}
	return a
}
