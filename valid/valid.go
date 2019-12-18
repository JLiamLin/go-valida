package valid

import (
	"github.com/go-playground/validator/v10"
	"valid/utils"
)

type Carrier interface {}

type Valid struct {
	Carrier			Carrier
	CustomFunc 		map[string]func(fl validator.FieldLevel) bool
}

// 将校验器绑定到给定的宿主上
func NewValid(carrier interface{}) (*Valid, error) {
	// 如果载体不是struct
	if err := utils.StructCheck(carrier); err != nil {
		return nil, err
	}
	// 初始化校验器
	v := &Valid{
		Carrier:    carrier,
		CustomFunc: CommonValidFunc(),
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

func (v *Valid) RegisterFunc(cf map[string]func(fl validator.FieldLevel) bool) *Valid {
	for key, val := range cf {
		v.CustomFunc[key] = val
	}
	return v
}

func (v *Valid) Register()  {
	
}

//func customFunc(fl validator.FieldLevel) bool {
//
//}

