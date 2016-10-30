package main

import "fmt"

func main() {

	// 字符串遍历
	str := "字符串遍历啊"
	l := len(str)

	for i := 0; i < l; i++ {
		ch := str[i]
		fmt.Println(ch)
	}

	for ch, i := range str {
		fmt.Println(ch, i)
	}
}
