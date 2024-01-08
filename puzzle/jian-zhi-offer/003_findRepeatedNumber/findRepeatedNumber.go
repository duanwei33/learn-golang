/*
题目描述
在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字是重复的，也不知道每个数字重复几次。请找出数组中任意一个重复的数字。

Input:
{2, 3, 1, 0, 2, 5}

Output:
2
*/

package main

import (
	"fmt"
)

func main() {
	test1 := []int{2, 3, 1, 0, 2, 5}
	test2 := []int{2, 3, 1, 0, 5, 5}
	fmt.Println(findRepeatedNumber(test1))
	fmt.Println(findRepeatedNumber(test2))

}

// 最简单的算法，两层循环，从第一个元素开始，和后面的元素一次比较，找到相同为止。
// 这个方法的时间复杂度是O(n^2)，其中n是切片的长度，因为使用了两层嵌套的循环。
// 这个方法的空间复杂度是O(1)，因为没有使用额外的空间。
func findRepeatedNumber1(numbers []int) (number int) {
	for index, num := range numbers {
		for i := index + 1; i <= len(numbers)-1; i++ {
			if num == numbers[i] {
				number = num
				return number
			}
		}
	}
	return -1
}

// 因为题目的特殊性：在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内
// 这个方法是构造一个数组，其长度有为n，其中原数组的中的数字是新数组中的索引，新数组中的元素就是出现的次数，初始值为0，当新数组元素=2时，这对应的原数组该数字重复了
// 这个方法的时间复杂度是O(n)，其中n是切片的长度
// 这个方法的空间复杂度是O(n),因为使用了一个与输入规模成线性关系的额外存储空间，即计数数组 s。在最坏的情况下，所有数字都是不同的，计数数组会与输入规模相同。
func findRepeatedNumber2(numbers []int) (number int) {
	s := make([]int, len(numbers))
	for i := 0; i <= len(numbers)-1; i++ {
		s[numbers[i]]++
		if s[numbers[i]] > 1 {
			return numbers[i]
		}
	}
	return -1
}

// 和findRepeatedNumber2类似，构造一个map，这样更具有普遍性
// 这个方法的时间复杂度是O(n)，其中n是切片的长度
// 这个方法的空间复杂度是O(n),因为使用了一个与输入规模成线性关系的map, 因为在最坏的情况下，所有数字都是不同的，map的大小将与输入规模相同。
func findRepeatedNumber3(numbers []int) (number int) {
	visitedNums := make(map[int]bool)
	for _, num := range numbers {
		if visitedNums[num] == true {
			return num
		}
		visitedNums[num] = true
	}
	return -1
}

// 因为题目的特殊性：在一个长度为 n 的数组里的所有数字都在 0 到 n-1 的范围内
// 可以采用原地置换法，即交换为止，使得这个数字的索引和值相等(nub[1]=1)，当发现该索引已有值时，则找到了重复的数字

func findRepeatedNumber(numbers []int) int {
	n := len(numbers)
	for i := 0; i < n; i++ {
		for numbers[i] != i {
			if numbers[i] == numbers[numbers[i]] {
				return numbers[i]
			}
			numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
		}
	}

	return -1 // 未找到重复数字
}
