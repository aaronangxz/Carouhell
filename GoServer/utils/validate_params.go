package utils

import (
	"fmt"
	"reflect"
)

//Validate if a is > b
func ValidateLimitMax(a uint32, b uint32) bool {
	return a > b
}

func ValidateUint(a *uint32) bool {
	return fmt.Sprint(reflect.TypeOf(a)) == "uint"
}

func ValidateString(a *string) bool {
	return fmt.Sprint(reflect.TypeOf(a)) == "string"
}
