package valid

import (
	"github.com/go-playground/validator/v10"
	"regexp"

	"valid/errors"
	"valid/utils"
)

const (
	ORIGINAL_CAUTION_STATUS 	= false
)

type Carrier interface {}

// 校验器
type Valid struct {
	// 宿主, 也即是要被绑定校验器的结构体
	Carrier			Carrier

	// 自定义的校验方法
	CustomFunc 		map[string]func(fl validator.FieldLevel) bool

	// 自定义的错误展示信息
	CustomMessage 	map[string]string

	// 展示具体错误信息状态值
	CautionStatus	bool
}

// 将校验器绑定到给定的宿主上
func NewValid(carrier interface{}) (*Valid, error) {
	// 如果载体不是struct
	if err := utils.StructCheck(carrier); err != nil {
		return nil, err
	}
	// 初始化校验器
	v := &Valid{
		Carrier:       		carrier,
		CustomFunc: 		CommonValidFunc(),
		CustomMessage:		CommonErrorMessage(),
		CautionStatus:		ORIGINAL_CAUTION_STATUS,
	}
	return v, nil
}

// 给宿主注入数据
func (v *Valid) Inject(data interface{}) *Valid {
	// 如果需要绑定的数据不是struct, 则直接跳过绑定
	if err := utils.StructCheck(data); err != nil {
		return v
	}
	// 获取宿主属性标签, 并从data中找出标签需要的的bind的数据, 绑定到宿主的属性上
	utils.Bind(data, v.Carrier)
	return v
}

// 注册自定义的校验方法, 如果用户定义的tag原本已存在, 则用户自定义的会覆盖原本存在的
func (v *Valid) RegisterValidFunc(rf map[string]func(fl validator.FieldLevel) bool) *Valid {
	for key, val := range rf {
		v.CustomFunc[key] = val
	}
	return v
}

// 注册自定义的校验错误显示, 如果用户定义的tag原本已存在, 则用户自定义的会覆盖原本存在的
func (v *Valid) RegisterErrorMessage(rm map[string]string) *Valid {
	for key, val := range rm {
		v.CustomMessage[key] = val
	}
	return v
}

// 关闭错误展示, 关闭错误展示后统一返回的错误信息为"非法请求参数"
func (v *Valid) CloseCaution() *Valid {
	v.CautionStatus = false
	return v
}

// 进行校验
func (v *Valid) Valid() error {
	vd := validator.New()
	// 如果存在CustomFunc, 则注册自定义校验
	if v.CustomFunc != nil {
		for tag, cf := range v.CustomFunc {
			_ = vd.RegisterValidation(tag, cf)
		}
	}

	var vf errors.ValidateFalseErrors
	// 进行校验
	if err := vd.Struct(v.Carrier); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fe := &errors.FieldValidateError{
				Field:    e.Field(),
				Rule:     e.Tag(),
				Param:    e.Param(),
				Describe: v.getDescribe(e),
			}
			vf = append(vf, fe)
		}
	}
	return vf
}

func (v *Valid) getDescribe(e validator.FieldError) string {
	// 全局禁止暴露错误描述
	if v.CautionStatus != false {
		return INVALID_PARAMS
	}
	// 获取标签定义的message信息, 粒度更小的禁止暴露错误描述
	tags, _ := utils.GetFieldsAndTags(v.Carrier, "")
	if msg, ok := tags[e.Field()]["message"]; ok {
		return msg
	}
	// 展示自定义的错误或者全局定义的错误
	describe := ""
	if msg, ok := v.CustomMessage[e.Field() + "~" + e.Tag()]; ok {
		describe = msg
	} else if msg, ok := v.CustomMessage[e.Tag()]; ok {
		describe = msg
	} else {
		return INVALID_PARAMS
	}
	// 格式化错误信息
	field := e.Field()
	if fd, ok := tags[e.Field()]["alias"]; ok {
		field = fd
	}
	fieldReg, _ := regexp.Compile("&{field}")
	describe = fieldReg.ReplaceAllString(describe, field)
	paramReg, _ := regexp.Compile("&{param}")
	describe = paramReg.ReplaceAllString(describe, e.Param())

	return describe
}



