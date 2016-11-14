package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

// 只有要修改本身值得时候才传递指针
func (a *Integer) AddRight(b Integer) {
	*a += b
}

func (a Integer) AddFalse(b Integer) {
	a += b
}

func main() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

	a.AddFalse(30)
	fmt.Println(a)

	a.AddRight(20)
	fmt.Println(a)

}
