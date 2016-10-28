package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// 没有条件的for是死循环，break或 return跳出
	for {
		fmt.Println("loop")
		break
	}
}
