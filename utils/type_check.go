package utils

import (
	"reflect"
	"valid/errors"
)

func StructCheck(it interface{}) error {
	val := reflect.ValueOf(it)
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return &errors.InvalidValidationError{Type: reflect.TypeOf(it)}
	}
	return nil
}

func GetFieldsAndTags(it interface{}, namespace string) (map[string]map[string]string, error) {
	if err := StructCheck(it); err != nil {
		return  nil, err
	}

	typeOfIt := reflect.TypeOf(it)
	if typeOfIt.Kind() == reflect.Ptr {
		typeOfIt = typeOfIt.Elem()
	}
	valueOfIt := reflect.ValueOf(it)
	if valueOfIt.Kind() == reflect.Ptr {
		valueOfIt = valueOfIt.Elem()
	}

	// 获取i的属性
	ft := make(map[string]map[string]string)
	for i := 0; i < typeOfIt.NumField(); i++ {
		fieldType := typeOfIt.Field(i)
		fieldValue := valueOfIt.Field(i)
		// 如果属性是结构体, 则递归处理
		if fieldType.Type.Kind() == reflect.Ptr && fieldValue.Type().Elem().Kind() == reflect.Struct {
			rft, err := GetFieldsAndTags(fieldValue.Interface(), fieldType.Name)
			if err != nil {
				return nil, err
			}
			for k, v := range rft {
				ft[k] = v
			}
		}
		// 获取不同的标签
		actualFieldName := fieldType.Name
		if namespace != "" {
			actualFieldName = namespace + "." + actualFieldName
		}
		tagMap := make(map[string]string)
		if alias := fieldType.Tag.Get("alias"); alias != "" {
			tagMap["alias"] = alias
		}
		if message := fieldType.Tag.Get("message"); message != "" {
			tagMap["message"] = message
		}
		if tagMap != nil {
			ft[actualFieldName] = tagMap
		}
	}

	return ft, nil
}
