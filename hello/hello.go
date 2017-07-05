package main

import "fmt"
import "unsafe"

const k int = 7

const (
	a2 = "abc"
	b2 = len(a2)
	c2 = unsafe.Sizeof(a2)
)

//在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1
const (
	a3 = iota
	b3 = iota
	c3 = iota
)

func main() {

	//第一种赋值方式
	var a int
	a = 1

	//第二种赋值方式
	var b = 2

	//第三种赋值方式
	c := 10

	var d, e int = 1, 2
	var f, g = 123, "hello"

	//这种不带声明格式的只能在函数体中出现
	h, i := 123, "hello"

	fmt.Println("Hello, 世界", a, b, c, d, e, f, g, h, i, k)

	main2()
}
