package main
// 利用test替换fmt包名，接下来可以用test来代替fmt使用
import test "fmt"

func main() {
	test.Println("测试")
	// 创建map的格式 make(map[key-type]value-type)
	m := make(map[string]int)

	// 赋值
	m["k1"] = 7
	m["k2"] = 13

	test.Println("map:", m)

	v1 := m["k1"]
	test.Println(v1)

	test.Println("len:" ,len(m))

	delete(m, "k2")

	test.Println(m)
	//如果map内该key存在_,_ 第一个返回值为其值，第二个为其是否存在
	//如果map内该key不存在_,_ 第一个返回值为0， 第二个为flase
	//简单来说，第二个是判断该key有没有存在于改map中的重要依据
	//格式为 ： value,key ：= map[key]
	prs, _ := m["k1"]
	_, prs1 := m["k1"]
	test.Println("prs: ",prs , " prs1 :", prs1)

	n := map[string]int{"foo": 1, "bar" :2}
	test.Println("map:", n)
}