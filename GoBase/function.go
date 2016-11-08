package main

import (
	"fmt"
)

func main() {

}

//无参无返回
func name() {
}

//无参有返回(单返回参数)
func name1() int {
}

//无参有返回(多返回参数)
func name2() (rec int, err error) {
}

//有参无返回 (多参数)
func name3(a, b int) {

}

// 不定参数(语法糖) 其相当于进去一个切片
func name4(args ...int) {
}
