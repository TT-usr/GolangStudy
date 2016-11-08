package main

import (
	"fmt"
)

func main() {
	var v1 int = 1
	var v2 int64 = 234
	var v3 string = ""
	var v4 float32 = 2.3
	name5(v1, v2, v3, v4)
}

//无参无返回
func name() {
}

//无参有返回(单返回参数)
func name1() int {
	return 1
}

//无参有返回(多返回参数)
func name2() (rec int, err error) {
	return 1, nil
}

//有参无返回 (多参数)
func name3(a, b int) {

}

// 不定参数(语法糖) 其相当于进去一个切片
// ...type 本质上是一个数组切片
func name4(args ...int) {

}

// 任意类型 不定参数用 ...interface{} 定义
func name5(args ...interface{}) {
	for _, ars := range args {
		switch ars.(type) {
		case int:
			fmt.Println(ars, "is a int value")
		case string:
			fmt.Println(ars, "is a string value")
		case int64:
			fmt.Println(ars, "is a int64 value")
		default:
			fmt.Println(ars, "is unknown type value")
		}
	}
}
