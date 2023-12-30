// 判断素数：
// 编写一个函数，判断给定的整数是否是素数。

package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := []int{1, 2, 3, 4, 11, 12, 12345}
	for _, nub := range numbers {
		if isPrime(nub) {
			fmt.Printf("The int %d is prime number.\n", nub)
		} else {
			fmt.Printf("The int %d is not prime number.\n", nub)
		}
	}
}

func isPrime(v int) (res bool) {
	res = true
	if v < 2 {
		fmt.Println("错误: 质数为大于1的正整数")
		res = false
	}
	// for i := v - 1; i > 2; i-- {
	for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
		if v%i == 0 {
			res = false
		}
	}
	return res
}
