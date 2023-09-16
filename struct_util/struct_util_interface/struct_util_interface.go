package struct_util_interface

import "reflect"

// CreateStructUtilT 工厂函数的type
type CreateStructUtilT func() StructUtilInterface

// StructUtilInterface 工具类的接口
type StructUtilInterface interface {

	// GetFields 获取结构体的所有属性，也就是field，可以传入结构体或结构体指针
	GetFields(subject interface{}) ([]reflect.StructField, error)

	// GetStructTag 获取结构体某个属性的某个标签值，可以传入结构体或结构体指针
	GetStructTag(subject interface{}, fieldName, tagName string) (string, error)

	// GetStructTags 获取结构体所有属性的某个标签值，返回属性名称到标签值的映射。可以传入结构体或结构体指针
	GetStructTags(subject interface{}, tagName string) (map[string]string, error)

	// GetFieldValue 获取结构体或结构体指针subject的fieldName属性值，可以传入结构体或结构体指针
	GetFieldValue(subject interface{}, fieldName string) (interface{}, error)

	// ModifyField 修改结构体某个属性的值，传入的必须是结构体指针
	// 把结构体指针subject里的fieldName属性修改为fieldValue，注意：如果fieldValue的数据类型与结构体属性不匹配的话，会panic
	ModifyField(subject interface{}, fieldName string, fieldValue interface{}) error
}
