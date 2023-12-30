// 反转字符串：
// 编写一个函数，接收一个字符串，返回其反转后的字符串。

package main

import (
	"fmt"
)

func main() {
	str := "abcde"
	fmt.Println(reverseString(str))
	fmt.Println(reverseStringRecursive(str))
	fmt.Println(reverseStringIterative(str))
}

// 使用[]byte+迭代实现反转字符串
func reverseString(str1 string) (str2 string) {
	str2Bytes := make([]byte, len(str1))
	for i := 0; i < len(str1); i++ {
		str2Bytes[len(str1)-1-i] = str1[i]
	}
	return string(str2Bytes)
}

// 使用递归实现反转字符串
func reverseStringRecursive(s string) string {
	if len(s) <= 1 {
		return s
	}
	return reverseStringRecursive(s[1:]) + string(s[0])
}

// 使用[]rune+迭代实现反转字符串
func reverseStringIterative(s string) string {
	sRunes := []rune(s)
	for i, j := 0, len(sRunes)-1; i < j; i, j = i+1, j-1 {
		sRunes[i], sRunes[j] = sRunes[j], sRunes[i]
	}
	return string(sRunes)
}
