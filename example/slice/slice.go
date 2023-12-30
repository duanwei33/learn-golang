package main

import "fmt"

func main() {
	// 定义切片方法一
	// var identifier []type: nil slice
	var s1 []string
	fmt.Printf("s1 set: %s\n", s1)
	fmt.Printf("The lenth of s1: %d\n", len(s1)) // 0

	// 定义切片方法二
	// var slice []type = make([]type, len): empyt slice
	// slice := make([]type, len)
	// make([]T, length, capacity)
	var s2 []string = make([]string, 2)
	fmt.Printf("s2 set: %s\n", s2)
	fmt.Printf("The lenth of s2: %d\n", len(s2))    // 2
	fmt.Printf("The capacity of s2: %d\n", cap(s2)) // 2

	s3 := make([]int, 3, 5)
	fmt.Printf("s3 set: %d\n", s3)
	fmt.Printf("The lenth of s3: %d\n", len(s3))    // 3
	fmt.Printf("The capacity of s3: %d\n", cap(s3)) // 5

}
