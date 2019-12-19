package valid

import (
	"github.com/go-playground/validator/v10"
	"valid/utils"
)

const (
	OriginalCautionStatus 	= true
)

type Carrier interface {}

// 校验器
type Valid struct {
	Carrier			Carrier
	CustomFunc 		map[string]func(fl validator.FieldLevel) bool
	CustomMessage 	map[string]string
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
		CautionStatus:		OriginalCautionStatus,
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

func (v *Valid) Valid() error {
	vd := validator.New()
	// 如果存在CustomFunc, 则注册自定义校验
	if v.CustomFunc != nil {
		for tag, cf := range v.CustomFunc {
			vd.RegisterValidation(tag, cf)
		}
	}
	// 进行校验
	if err := vd.Struct(v.Carrier); err != nil {
		// TODO错误信息处理
	}
}


