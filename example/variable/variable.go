package main

import "fmt"

func main() {
	// 如果我们使用一个变量给另外一个变量赋值，那么真正赋给后者的，并不是前者持有的那个值，而是该值的一个副本。
	v1 := "HelloV1"
	v2 := v1
	v1 = "NewHelloV1"
	fmt.Println(v2) // HelloV1

}
