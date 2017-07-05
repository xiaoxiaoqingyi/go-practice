package main

import "fmt"

const (
	i1 = 1 << iota
	j1 = 3 << iota
	k1
	l1
)

const (
	kk = iota + 2
	jj
)

const (
	dd = iota
	ff = 2
	aa = iota
)

func main2() {

	fmt.Println("i=", i1)
	fmt.Println("j=", j1)
	fmt.Println("k=", k1)
	fmt.Println("l=", l1)
	fmt.Println("kk=", kk)
	fmt.Println("jj=", jj)
	fmt.Println("dd=", dd)
	fmt.Println("ff=", ff)
	fmt.Println("aa=", aa)
}
