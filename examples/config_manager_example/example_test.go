package config_manager_example

import (
	"fmt"
	"go-utils/config_manager"
	"testing"
)

// 程序运行时的当前目录，等于程序启动文件所在的目录
func Test_Example(t *testing.T) {
	cm, _ := config_manager.CreateConfigManager("yaml", []string{"./config"}, "yaml_config")
	fmt.Println(cm.Get("Mysql"))
	fmt.Println(cm.GetInt("Mysql.port"))
	fmt.Println(cm.GetString("Mysql.UserName"))
	fmt.Println(cm.GetBool("Mysql.Writable"))
	fmt.Println(cm.GetDuration("Mysql.Timeout").Minutes())
	// 输出内容：
	// map[port:3106 username:root writable:false]
	// 3106
	// root
	// false
	// 30

	cm, _ = config_manager.CreateConfigManager("json", []string{"./config"}, "json_config")
	fmt.Println(cm.Get("student"))
	fmt.Println(cm.GetInt("student.age"))
	fmt.Println(cm.GetString("student.name"))
	fmt.Println(cm.GetStringSlice("schools"))
	// 输出内容：
	// map[age:22 name:luan]
	// 22
	// luan
	// [collageA collageB]

}
