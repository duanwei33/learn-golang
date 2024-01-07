package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 设置随机数种子，可以使用当前时间作为种子

	// 示例字符串切片
	strSlice := []string{"apple", "banana", "orange", "grape"}

	// 随机取一个元素
	randomElement := getRandomElementString(strSlice)
	fmt.Println("随机取得的元素:", randomElement)
	fmt.Println("随机取得的数字:", getRandomNum(2, 8))
}

func getRandomElementString(strSlice []string) string {
	// 生成一个在切片范围内的随机索引
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(strSlice))
	fmt.Println(randomIndex)
	// 返回随机索引对应的元素
	return strSlice[randomIndex]
}

func getRandomNum(m int64, n int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(n-m+1) + m
}
