package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 3, 8, 1, 6}
	// 外层控制这次冒泡第几大的数 i不是指针，只是计数用
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ { // 因为最右侧的第n大已经排序好了，就不用继续排了
			if arr[j] > arr[j+1] {
				// 交换当前和后一个
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
