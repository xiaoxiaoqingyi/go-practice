package main

import "fmt"

func main() {
	const (
		a = iota + 1 //0
		b = 100      //1
		c            //2
		d = "ha"     //独立值，iota += 1
		e
		e1       //"ha"   iota += 1
		f  = 100 //iota +=1
		g
		g1        //100  iota +=1
		h  = iota //7,恢复计数
		i         //8
	)
	fmt.Println(a, b, c, d, e, e1, f, g, g1, h, i)
}
