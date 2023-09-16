package struct_util

import (
	"errors"
	"go-utils/struct_util/struct_util_interface"
	"reflect"
)

type structUtil struct {
}

var s = &structUtil{}

func CreateStructUtil() struct_util_interface.StructUtilInterface {
	return s
}

// GetFields 获取结构体的所有属性，也就是field，可以传入结构体或结构体指针
func (su *structUtil) GetFields(subject interface{}) ([]reflect.StructField, error) {
	fields := make([]reflect.StructField, 0)

	typeOfSubject := reflect.TypeOf(subject)

	//switch里面判断subject的类型，如果是结构体指针类型则做一系列转换，获取结构体类型
	switch typeOfSubject.Kind() {
	case reflect.Struct:
		break
	case reflect.Ptr: //如果是指针类型，则需要通过Elem()函数得到它的实际数据类型
		for typeOfSubject.Kind() == reflect.Ptr {
			typeOfSubject = typeOfSubject.Elem()
		}
		if typeOfSubject.Kind() != reflect.Struct { //如果实际数据类型不是结构体类型，则返回错误
			return fields, errors.New("error: subject can not be " + typeOfSubject.Kind().String())
		}
	default: //如果不是结构体类型也不是指针类型，则返回错误
		return fields, errors.New("error: subject can not be " + typeOfSubject.Kind().String())
	}

	for i := 0; i < typeOfSubject.NumField(); i++ {
		fields = append(fields, typeOfSubject.Field(i))
	}
	return fields, nil
}

// GetStructTag 获取结构体某个属性的某个标签值，可以传入结构体或结构体指针
func (su *structUtil) GetStructTag(subject interface{}, fieldName, tagName string) (string, error) {
	if subject == nil {
		return "", errors.New("error: subject can not be nil")
	}

	typeOfSubject := reflect.TypeOf(subject)

	//switch里面判断subject的类型，如果是结构体指针类型则做一系列转换，获取结构体类型
	switch typeOfSubject.Kind() {
	case reflect.Struct:
		break
	case reflect.Ptr: //如果是指针类型，则需要通过Elem()函数得到它的实际数据类型
		for typeOfSubject.Kind() == reflect.Ptr {
			typeOfSubject = typeOfSubject.Elem()
		}
		if typeOfSubject.Kind() != reflect.Struct { //如果实际数据类型不是结构体类型，则返回错误
			return "", errors.New("error: subject can not be " + typeOfSubject.Kind().String())
		}
	default: //如果不是结构体类型也不是指针类型，则返回错误
		return "", errors.New("error: subject can not be " + typeOfSubject.Kind().String())
	}

	if field, ok := typeOfSubject.FieldByName(fieldName); ok {
		return field.Tag.Get(tagName), nil
	} else {
		return "", errors.New("error: subject doesn't has the field: " + fieldName)
	}
}

// GetStructTags 获取结构体所有属性的某个标签值，返回属性名称到标签值的映射。可以传入结构体或结构体指针
func (su *structUtil) GetStructTags(subject interface{}, tagName string) (map[string]string, error) {
	if fields, err := su.GetFields(subject); err == nil {
		r := make(map[string]string)
		for _, v := range fields {
			r[v.Name] = v.Tag.Get(tagName)
		}
		return r, nil
	} else {
		return nil, err
	}

}

// GetFieldValue 获取结构体或结构体指针subject的fieldName属性值，可以传入结构体或结构体指针
func (su *structUtil) GetFieldValue(subject interface{}, fieldName string) (interface{}, error) {
	valueOfSubject := reflect.ValueOf(subject)

	// 老规矩，还是先判断一下传入参数的数据类型，如果是指针则进行取值处理
	if valueOfSubject.Kind() == reflect.Ptr {
		for valueOfSubject.Kind() == reflect.Ptr {
			valueOfSubject = valueOfSubject.Elem()
		}
	}
	if valueOfSubject.Kind() != reflect.Struct {
		return nil, errors.New("subject is not a pointer of struct or struct")
	}

	// 如果该属性存在的话，field不是零值
	if field := valueOfSubject.FieldByName(fieldName); !field.IsZero() {
		return field.Interface(), nil
	} else {
		// 如果属性不存在，则直接返回错误
		return nil, errors.New("field: " + fieldName + " not exist in subject")
	}

}

// ModifyField 修改结构体某个属性的值，传入的必须是结构体指针
// 把结构体指针subject里的fieldName属性修改为fieldValue，注意：如果fieldValue的数据类型与结构体属性不匹配的话，会panic
func (su *structUtil) ModifyField(subject interface{}, fieldName string, fieldValue interface{}) error {

	valueOfSubject := reflect.ValueOf(subject)

	// 老规矩，还是先判断一下传入参数的数据类型，如果是指针则进行取值处理
	if valueOfSubject.Kind() == reflect.Ptr {
		for valueOfSubject.Kind() == reflect.Ptr {
			valueOfSubject = valueOfSubject.Elem()
		}
	} else {
		return errors.New("subject is not a pointer")
	}
	if valueOfSubject.Kind() != reflect.Struct {
		return errors.New("subject is not a pointer of struct")
	}

	// 先获取要修改的属性，如果该属性存在的话，field不是零值
	if field := valueOfSubject.FieldByName(fieldName); !field.IsZero() {
		if field.CanSet() {
			// 有一些属性是不能够从外部修改的，比如私有属性，所以先判断一下能不能修改
			field.Set(reflect.ValueOf(fieldValue))
			return nil
		} else {
			return errors.New("field: " + fieldName + " in subject can not be set")
		}

	} else {
		// 如果属性不存在，则直接返回错误
		return errors.New("field: " + fieldName + " not exist in subject")
	}
}
