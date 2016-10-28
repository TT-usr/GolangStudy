package main

import "fmt"

type Vertex struct {
	x, y int
	z    string
}

type Ver struct {
	x, y int
}

var (
	v1 = Ver{1, 2}  // 类型为 Ver
	v2 = Ver{X: 1}  // Y:0 被省略
	v3 = Ver{}      // X:0 和 Y:0
	v4 = &Ver{1, 2} // 类型为 *Ver
)

func main() {
	fmt.Println(Vertex{1, 2, "测试"})

	v := Vertex{1, 2, "干嘛"}
	v.x = 4
	v.y = 12
	fmt.Println(v)

	p := &v
	p.x = 100
	fmt.Println(p)

	fmt.Println(v1, v2, v3, v4)
}
