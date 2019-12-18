package utils

import (
	"reflect"
)

// 从dst结构体的属性中的标签中获取bind的值
// 然后从src中找到bind的值所对应的字段值, 如果类型与dst中的类型一致, 则赋值给dst对应的属性
// ================= example:
// Dst的类型
// type Dst struct {
//	Name 		string		`bind:"Name"`
//	Age 		int
//	Hobby 		[]string	`bind:"Hobby"`
//	Friend 		*FriendDst	`bind:"ID"`
// }
// type FriendDst struct {
//	Name		string	`bind:"Name"`
//	Gender 		int		`bind:"Gender"`
// }
//
// Src的类型
// type Src struct {
//	Name 		string
//	Age 		int
//	Hobby 		[]string
//	Friend 		*FriendSrc
// }
// type FriendSrc struct {
//	Name		string
//	Gender 		int
// }
//
// src的数据
// friendSrc := &FriendSrc{
//	Name:   	"林先生",
//	Gender:     1,
// }
//	src := &Src {
//	Name:    	"Liam",
//	Age:     	25,
//	Hobby:   	[]string{"睡觉", "学习"},
//	Friend:     friendSrc,
// }
//
// 初始化dst
// dst := &Dst{Friend:&FriendDst{}}
//
// utils.Bind(src, dst)
// 得到的dst的值为
// &Dst {
//	Name:		"Liam"
//	Age:		0 // 初始化的值, 因为tag里面没有指明要bind的值
// 	Hobby:		[]string{"睡觉", "学习"},
// 	Friend:		&FriendDst {
//		Name:		"林先生",
//		Gender:		1,
// 	   }
// }
// ================= 注意事项:
// 1. 该方法默认为安全调用, 即默认传参src, dst均为struct类型或者struct类型的指针, 不再做传参的类型判断, 参数类型错误会 panic
// 2. 方法使用递归来支持多层级嵌套struct, 但是嵌套的struct需要初始化, 例如 dst := &Dst{Friend:&FriendDst{}}, 如果使用
//	  dst := &Dst{} 则Friend没有赋值, 因为dst.Friend为nil, 取不到dst.Friend的属性
func Bind(src, dst interface{}) {
	// 如果src为空, 直接返回
	valueOfSrc := reflect.ValueOf(src)
	if valueOfSrc.IsNil() {
		return
	}
	if valueOfSrc.Kind() == reflect.Ptr {
		valueOfSrc = valueOfSrc.Elem()
	}

	// 获取dst类型和值信息
	valueOfDst := reflect.ValueOf(dst)
	if valueOfDst.IsNil() {
		return
	}
	if valueOfDst.Kind() == reflect.Ptr {
		valueOfDst = valueOfDst.Elem()
	}
	typeOfDst := reflect.TypeOf(dst)
	if typeOfDst.Kind() == reflect.Ptr {
		typeOfDst = typeOfDst.Elem()
	}

	// 获取dst的结构属性
	for i := 0; i < typeOfDst.NumField(); i++ {
		// 获取bind标签值
		dstFieldType := typeOfDst.Field(i)
		bindTag := dstFieldType.Tag.Get("bind")
		if bindTag == "" {
			continue
		}
		dstFieldValue := valueOfDst.Field(i)

		// 如果在src没找到对应的bind属性, 则跳过
		srcFieldValue := valueOfSrc.FieldByName(bindTag)
		if !srcFieldValue.IsValid() {
			continue
		}
		// 如果属性是结构体, 则递归处理
		if dstFieldType.Type.Kind() == reflect.Ptr && dstFieldValue.Type().Elem().Kind() == reflect.Struct {
			Bind(srcFieldValue.Interface(), dstFieldValue.Interface())
		}
		// 类型相同, 则进行赋值
		if srcFieldValue.Type() == dstFieldValue.Type() && dstFieldValue.CanSet() {
			dstFieldValue.Set(srcFieldValue)
		}
	}
}
