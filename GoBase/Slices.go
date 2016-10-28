package main

import "fmt"

func main() {
	// slices 切片 定义格式make([]类型, 初始化长度)
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	// len 长度
	fmt.Println("len:", len(s))

	var z []int
	fmt.Println(z, len(z))
	if z == nil {
		fmt.Println("nil!")
	}
	// 拼接
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// 拷贝
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 分割(截取2 ~ 5)
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 截取到5
	l = s[:5]
	fmt.Println("sl2:", l)

	// 	从2开始截取
	l = s[2:]
	fmt.Println("sl3", l)

	// 另一种初始化
	t := []string{"g", "h", "i"}
	fmt.Println(t)

	// 	二维切片
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
