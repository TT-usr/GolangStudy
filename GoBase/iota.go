package main

import "fmt"

// 在下一个iota出现之前，iota其所代表的数字自动++
// iota 会在每个const开头被重设为0
const (
	c0 = iota //0
	c1 = iota //1
	c2 = iota //2
)

const (
	cc0 = iota //0
	cc1 = iota //1
)

//可以简写为
const (
	a = iota
	b
	c
)

func main() {
	fmt.Println(c0, c1, c2, cc0, cc1)
	fmt.Println("测试一下", c0, "占位符") //目前测试Go不需要占位符
}
