package struct_util

import (
	"fmt"
	"go-utils/struct_util"
	"testing"
)

// User 随便声明一个User结构体，用于测试
type User struct {
	Name string `json:"name" gorm:"user_name"`
	Age  int    `json:"age" gorm:"user_age"`
}

// 先创建一个工具类对象
var su = struct_util.CreateStructUtil()

// Test1 GetStructTag方法的使用样例
func Test1(t *testing.T) {
	u := User{}
	// 测试样例1，控制台输出 user_name
	if tagValue, err := su.GetStructTag(u, "Name", "gorm"); err == nil {
		fmt.Println(tagValue)
	} else {
		fmt.Println(err.Error())
	}

	// 测试样例2，控制台输出 age
	if tagValue, err := su.GetStructTag(u, "Age", "json"); err == nil {
		fmt.Println(tagValue)
	} else {
		fmt.Println(err.Error())
	}
}

// Test2 GetStructTags方法的使用样例
func Test2(t *testing.T) {
	u := User{}
	// 测试样例1，控制台输出 map[Age:user_age Name:user_name]
	if tagValues, err := su.GetStructTags(u, "gorm"); err == nil {
		fmt.Println(tagValues)
	} else {
		fmt.Println(err.Error())
	}

	// 测试样例2，控制台输出 map[Age: Name:]
	if tagValues, err := su.GetStructTags(u, "nooooo"); err == nil {
		fmt.Println(tagValues)
	} else {
		fmt.Println(err.Error())
	}
}

// 随便声明的结构体，用于测试
type Student struct {
	Name    string `json:"name" gorm:"user_name"`
	Age     int    `json:"age" gorm:"user_age"`
	Friends []Student
}

// Test3 GetFields方法的使用样例
func Test3(t *testing.T) {
	// 测试样例1，输出内容如下：
	// 属性名: Name, 数据类型名称: string
	// 属性名: Age, 数据类型名称: int
	// 属性名: Friends, 数据类型名称: []struct_util.Student
	if fields, err := su.GetFields(&Student{}); err == nil {
		for _, f := range fields {
			fmt.Printf("属性名: %s, 数据类型名称: %s\n", f.Name, f.Type.String())
		}
	} else {
		fmt.Println(err.Error())
	}
}

// Person Car 随便声明两个具有嵌套关系的结构体，用于测试
type Person struct {
	Name        string
	age         int
	PersonalCar *Car
}

func (p *Person) showInfo() {
	fmt.Printf("my name is %s, I am %d years old. My car is %s and worth %d$\n",
		p.Name, p.age, p.PersonalCar.Brand, p.PersonalCar.Price)
}

type Car struct {
	Brand string
	Price int
}

// Test4 ModifyField、PersonalCar方法的使用样例
func Test4(t *testing.T) {
	p := Person{
		Name: "张三",
		age:  18,
		PersonalCar: &Car{
			Brand: "奥迪",
			Price: 500,
		},
	}

	// 成功修改属性Name，控制台输出：my name is 李四, I am 18 years old. My car is 奥迪 and worth 500$
	if err := su.ModifyField(&p, "Name", "李四"); err == nil {
		p.showInfo()
	} else {
		fmt.Println(err.Error())
	}

	// 不能修改age属性，因为age是私有属性，控制台输出：field: age in subject can not be set
	if err := su.ModifyField(&p, "age", 99); err == nil {
		p.showInfo()
	} else {
		fmt.Println(err.Error())
	}

	// 能正常修改PersonalCar属性，控制台输出：my name is 李四, I am 18 years old. My car is 五菱 and worth 900$
	if err := su.ModifyField(&p, "PersonalCar", &Car{Brand: "五菱", Price: 900}); err == nil {
		p.showInfo()
	} else {
		fmt.Println(err.Error())
	}

	// 能正确获取PersonalCar属性，控制台输出 &{五菱 900}
	if personCar, err := su.GetFieldValue(p, "PersonalCar"); err == nil {
		fmt.Println(personCar)
	} else {
		fmt.Println(err.Error())
	}

	// "car"是字符串变量，不能赋值给*Car类型的PersonalCar，会出现panic
	if err := su.ModifyField(&p, "PersonalCar", "car"); err == nil {
		p.showInfo()
	} else {
		fmt.Println(err.Error())
	}

}
