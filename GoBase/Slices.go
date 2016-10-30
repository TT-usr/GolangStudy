package main

import "fmt"

func main() {

	creatSlice()

	bianliSlice()

	calucateSlice()

	appendSlice()

	copySlice()

	cutSlice()

	twoSlice()
}

func creatSlice() {
	//定义数组
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 基于数组创建一个数组切片
	var mySlice []int = myArray[:5]

	fmt.Println(mySlice)

	// 额外的初始化方法
	// slices 切片 定义格式make([]类型, 初始化长度)
	// mySlice3 = make([]int 3)

	// 初始个数为3 并预留5个元素空间
	//mySlice1 = make([]int, 3, 5)

	// 包含5个元素的数组切片
	//mySlice2 := []int{1, 2, 3, 4, 5}

	// modifyArr(mySlice)
	// fmt.Println(mySlice)
}

func bianliSlice() {
	var mySlice [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(mySlice); i++ {
		fmt.Println("mySlice[", i, "] = ", mySlice[i])
	}
	// 或者使用range来遍历
	for _, v := range mySlice {
		fmt.Println(v)
	}

}

func calucateSlice() {
	var mySlice [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l := len(mySlice)
	s := cap(mySlice)
	fmt.Print(l, "  ", s)

}

func appendSlice() {
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
	// 或者
	s = append(s, []string{"1", "2", "3"}...)
	// 这种拼接必须加三个点，如果不加就会爆语法错误
	// 补充一下， 如果切片容量不足，内部实现为：系统会自动配一块足够大的内存，将切片复制过来
	//（尽量在你已知的范围内给其分配足够大的空间，以减小CPU损耗，
	// 因为系统并不知道足够大是多大，，分配和复制可能会进行很多次

}

func cutSlice() {
	s := make([]string, 6)
	// 分割(截取2 ~ 5)
	l := s[2:5] //这要这个取值不大于s的len即可。就不会产生异常
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

}

func copySlice() {
	// 或者
	s1 := []string{"1", "2", "3"}

	// 拷贝
	c := make([]string, len(s1))
	copy(c, s1)
	fmt.Println("cpy:", c)
}

func twoSlice() {
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

func modifySlice(slice []int) {
	// slice[0] = "789"
	// fmt.Print(slice, "")

}
