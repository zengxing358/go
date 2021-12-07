package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//排序
	var nums = []int{5, 1, 1, 2, 0, 0}
	fmt.Println(sortArray(nums))
	//打乱
	var arr = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	fmt.Println(shuffle(arr))
	//转换大小写
	var tmp string = "您好ABZaz"
	fmt.Println(transe(tmp))
}
func transe(arr string) string {
	aBytes := []rune(arr)
	var ans string = ""
	for i := 0; i < len(aBytes); i++ {
		if aBytes[i] >= 65 && aBytes[i] <= 90 {
			ans += string(aBytes[i] + 32)
		} else if aBytes[i] >= 97 && aBytes[i] <= 122 {
			ans += string(aBytes[i] - 32)
		} else {
			ans += string(aBytes[i])
		}
	}
	return ans
}
func shuffle(arr []string) []string {
	n := len(arr)
	var ans []string
	var p int
	for i := 0; i < n; i++ {
		p = rand.Intn(len(arr))
		ans = append(ans, arr[p])
		arr = append(arr[:p], arr[p+1:]...)
	}
	return ans
}
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) >> 1
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	var temp []int
	i := 0
	j := 0
	for {
		if i >= len(left) || j >= len(right) {
			break
		}
		if left[i] <= right[j] {
			temp = append(temp, left[i])
			i += 1
		} else {
			temp = append(temp, right[j])
			j += 1
		}
	}
	temp = append(temp, left[i:]...)
	temp = append(temp, right[j:]...)
	return temp
}
