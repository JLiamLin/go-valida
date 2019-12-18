package errors

import "reflect"

type InvalidValidationError struct {
	Type reflect.Type
}

func (e *InvalidValidationError) Error() string {

	if e.Type == nil {
		return "错误的宿主类型, 类型为 (nil)"
	}

	return "错误的宿主类型, 类型为 ( " + e.Type.String() + ")"
}

