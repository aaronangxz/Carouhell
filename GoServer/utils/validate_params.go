package utils

import (
	"fmt"
	"reflect"

	"github.com/aaronangxz/TIC2601/models"
)

//Validate if a is > b
func ValidateLimitMax(a uint32, b uint32) bool {
	return a > b
}

func ValidateMaxStringLength(s string) bool {
	return len(s) < int(models.MaxStringLength)
}

func ValidateUint(a *uint32) bool {
	return *a >= uint32(0)
}

func ValidateString(a *string) bool {
	return fmt.Sprint(reflect.TypeOf(a)) == "*string"
}
