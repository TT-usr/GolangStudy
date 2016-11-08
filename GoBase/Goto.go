package main

import (
	"fmt"
)

func main() {
	myFunc()
}

func myFunc() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE //跳转标签
	}
}
