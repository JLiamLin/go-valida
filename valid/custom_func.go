package valid

import "github.com/go-playground/validator/v10"

// 公共的自定义校验规则
func CommonValidFunc() map[string]func(fl validator.FieldLevel) bool {
	vf := make(map[string]func(fl validator.FieldLevel) bool)

	// 自定义 commonExample 的校验规则 (示例)
	vf["commonExample"] = func(fl validator.FieldLevel) bool {
		return fl.Field().String() == fl.Param()
	}

	return vf
}
