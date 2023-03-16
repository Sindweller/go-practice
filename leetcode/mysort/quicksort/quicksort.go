package main

import "fmt"

// 每次partition都能将一个位置放到最终位置上
func main() {
	arr := []int{3, 10, 4, 2, 5, 333, 1, 3, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func quickSort(arr []int, left, right int) {
	if left < right {
		// 确定p的位置
		p := partition(arr, left, right)
		// 左右子数组划分
		quickSort(arr, left, p-1)
		quickSort(arr, p+1, right)
	}
}

func partition(arr []int, left, right int) int {
	pivot := arr[right]
	fmt.Println(pivot)
	// 需要交替移动
	for left < right {
		// 不需要交换
		for left < right && arr[left] <= pivot {
			left++
		}
		// 这里left> right的值了
		// 交换
		arr[right] = arr[left]
		// 该从右往左找了
		for left < right && arr[right] >= pivot {
			right--
		}
		// 这里right < pivot了 跟left交换
		arr[left] = arr[right]
	}
	fmt.Println(arr)
	// 填充空位
	arr[right] = pivot
	return right
}
