package utils

import (
	"reflect"
	"valid/errors"
)

func StructCheck(uc interface{}) error {
	val := reflect.ValueOf(uc)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return &errors.InvalidValidationError{Type: reflect.TypeOf(uc)}
	}
	return nil
}
