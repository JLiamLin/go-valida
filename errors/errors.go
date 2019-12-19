package errors

import (
	"bytes"
	"reflect"
	"strings"
)

type InvalidValidationError struct {
	Type reflect.Type
}

func (e *InvalidValidationError) Error() string {

	if e.Type == nil {
		return "错误的宿主类型, 类型为 (nil)"
	}

	return "错误的宿主类型, 类型为 ( " + e.Type.String() + ")"
}

type FieldValidateError struct {
	// 校验的字段
	Field 			string

	// 校验的规则
	Rule			string

	// 规则对应的参数
	Param 			string

	// 错误描述
	Describe 		string
}

func (fe *FieldValidateError) Error() string {
	return fe.Describe
}

type ValidateFalseErrors []*FieldValidateError

func (ve ValidateFalseErrors) Error() string {

	buff := bytes.NewBufferString("")

	var fe *FieldValidateError

	for i := 0; i < len(ve); i++ {

		fe = ve[i]
		buff.WriteString(fe.Error())
		buff.WriteString(";")
	}

	return strings.TrimSpace(buff.String())
}

