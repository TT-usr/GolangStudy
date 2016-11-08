package main

import (
	"fmt"
)

func main() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Printf("这里拿到了错误 -> %d\n", err) //这里的err其实就是panic传入的内容，55
		}
		fmt.Println("d")
	}()
	f()
}

func f() {
	fmt.Println("a")
	panic(55) //这里抛出错误
	fmt.Println("b")

	fmt.Println("f")
}
