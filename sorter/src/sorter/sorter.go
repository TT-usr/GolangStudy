package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"algorithm/bubblesort"
	"algorithm/qsort"
)

// 参数1：定义使用标识
// 参数2：值得名字
// 参数3：该值得描述
var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

// flag.String(name, value, usage)
// 或者使用以下方式定义：
// var algorithm *string
// flag.StringVar(*algorithm, "a", "qsort", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile) //打开文件
	if err != nil {              //如果打不开
		fmt.Println("Failed to open the input file ", infile)
		return
	}

	defer func() { // 加了完成后调用闭包
		file.Close() //保证以下代码出错 均可以正确的关闭文件
		fmt.Println("file close")
	}()

	br := bufio.NewReader(file)
	values = make([]int, 0) // 初始化一个int类型的切片

	for { //死循环

		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF { //文件到达了结尾
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unpected")
			return
		}

		str := string(line) //转换字符数组为字符串

		value, err1 := strconv.Atoi(str) //将字符串转换为十进制整数 ParseInt(s, 10, 0) 的简写

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value) // 拼接数据
	}
	return
}

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Fail to creat the output file ", outfile)
		return err
	}

	defer func() {
		file.Close()
	}()

	for _, value := range values {
		str := strconv.Itoa(value) //将整数转换为十进制字符串形式（即：FormatInt(i, 10) 的简写）
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {

	flag.Parse()

	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =", *algorithm)
	}

	values, err := readValues(*infile)

	if err == nil {

		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.Bubblesort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}

		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.") //截取两次时间的差值

		writeValues(values, *outfile)

	} else {
		fmt.Println(err)
	}
}
