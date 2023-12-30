// 找最大值：
// 编写一个函数，接收一个整数切片，返回切片中的最大值。

package main

import "fmt"

func main() {
	s := []int{1, 2, 5, 12, 11, 14}
	maxValue := findMaxValue(s)
	fmt.Println("Max Value:", maxValue)
}

func findMaxValue(s []int) (max int) {
	if len(s) == 0 {
		// 切片为空时，返回零或其他合适的默认值
		fmt.Println("切片为空，请检查！")
		return 0
	}

	max = s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}
