package main

import "fmt"

func main() {
	// 这样初始化  每个元素默认初始值为0
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("set:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	// 二维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// 遍历数组
	for _, v := range a {
		fmt.Println(v)
	}

	modifyArr(b)
	fmt.Println(b)
}

// 传递的不是指针， 内部无法修改外部传进来的数组变量
func modifyArr(array [5]int) {
	array[0] = 100
	fmt.Println(array)
}
