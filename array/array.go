package main

import "fmt"

func main() {

     var balance = [5]float32{1000.0, 11.0, 12.3, 22.3, 44.2}

     //如果忽略 [] 中的数字，不设数组大小，Go 语言 会根据元素的个数来设置数组大小：
     var balance2 = [...] float32{1000.0, 11.0, 12.3, 22.3, 44.2}
	

     for i := 0; i < len(balance); i++ {
          fmt.Printf("balance[%d] = %d \n", i, balance2[i])
     }
}
