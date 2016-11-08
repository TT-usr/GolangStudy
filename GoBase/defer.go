package main

import (
	"fmt"
)

// defer主要用于不管任何情况下，都能保证会被执行
func main() {
	// defer 语句会延迟函数的执行直到上层函数返回。
	// defer 只能描述方法？！

	// 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。

	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// 打印结果 ： 9876543210
		// 原因：延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用
		defer fmt.Print(i)
	}

	fmt.Println("done")

	defer func() {
		//延迟闭包，如果觉得defer描述一条语句不够，可用闭包来解决
	}()

}
