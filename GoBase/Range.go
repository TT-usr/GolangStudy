package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// range相当于遍历数组与切片
	nums := []int{2, 3, 4}
	sum := 0
	// num 相当于一个临时变量，把每一次从nums取得值赋给num
	// _ 代表值忽略，这里代表着不取key
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	// 字符串用%s来格式化
	kvs := map[string]string{"a": "apple", "b": "banana"}
	// 这里的k, v 两个临时变量对应着从kvs里拿出key 和 value 分别赋值给k , v
	for k, v := range kvs {
		fmt.Println("%s -> %s", k, v)
	}
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
