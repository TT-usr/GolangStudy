package main

import "fmt"

func main() {
	// 类型为小写
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// 只有初始化才用 := 赋值。
	// a := 9 等价于  var a int = 9
	f := "short"
	fmt.Println(f)
}
