# go-utils

## 概述
本仓库主要是对一些标准库和一些流行的第三方库进行二次封装，从而实现一些简单易用且高效的工具类和工具方法。

## 工具类

### configManager

在后端开发中，我们经常会使用配置文件，所以我封装出这样一个工具类出来，能够提供简洁的接口，让我们很方便地读取配置文件并从配置文件中提取信息。

工具类的名称是<code>configManager</code>，主要有以下功能：
- 根据配置文件的路径和文件名读取配置信息
- 通过一系列Get函数，可以根据key查询配置项的值
- 自动将配置项写入缓存，提高配置项查询速率
- 自动监听配置文件变化，自动更新配置项查询结果，可以添加一个或多个钩子函数处理文件变化事件

工具类实现代码位于[/config_manager](./config_manager)，使用样例代码位于[/examples/config_manager_example](./examples/config_manager_example)

### search

提供了在切片中查询特定元素的方法，主要功能如下：
- 从一个无序列表中查找一个、或多个指定元素（时间复杂度O(n)，假如有个元素，下同）
- 从有序列表中二分查找指定元素（时间复杂度O(logn)）
- 查找最值元素，最值的定义由用户传入的比较函数决定，可以是最大值、最小值、最大的绝对值、最长的字符串等等（时间复杂度O(n*m)，m是比较函数的时间复杂度）
- 使用泛型技术，支持各种数据类型的查找

1.18泛型不支持结构体方法，所以就直接写成普通函数的形式，工具函数实现代码位于[/search](./search)，使用样例代码位于[/examples/search_example](./examples/search_example)


### structUtil

工具类的名称是<code>configManager</code>，主要有以下功能：
- 获取结构体某属性或所有属性的标签值
- 获取结构体所有属性的详细信息，包括数据类型、属性名等等
- 根据属性名读取或修改结构体的属性值

工具类实现代码位于[/struct_util](./struct_util)，使用样例代码位于[/examples/struct_util](./examples/struct_util)

