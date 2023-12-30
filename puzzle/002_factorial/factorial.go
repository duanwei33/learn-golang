// 阶乘计算：
// 编写一个函数，计算给定正整数的阶乘。

package main

import "fmt"

func main() {
	fmt.Println(factorialWithFor(4)) // 24
	fmt.Println(factorialWithFor(0)) // "You should input a value > 0" 0

	fmt.Println(factorialWithRecursion(4)) // 24
	fmt.Println(factorialWithRecursion(0)) // "You should input a value > 0" 0

}

func factorialWithFor(v int) (res int) {
	if v < 1 {
		fmt.Println("You should input a value > 0")
		return 0
	}
	for res = 1; v > 1; v-- {
		res = v * res
	}
	return res
}

func factorialWithRecursion(v int) (res int) {
	if v < 1 {
		fmt.Println("You should input a value > 0")
		return 0
	}

	if v == 1 {
		return 1
	}
	return v * factorialWithRecursion(v-1)
}
