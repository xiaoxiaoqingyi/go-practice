package main

import "fmt"

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

	fmt.Println("Hello, 世界", a, b, c, d, e, f, g, h, i)
}
