package main

import "fmt"

// 每次partition都能将一个位置放到最终位置上
func main() {
	arr := []int{3, 10, 4, 2, 5, 333, 1, 3, 2}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low, high int) {
	if low < high {
		p := partition(arr, low, high) // 找到枢轴的位置
		// 左右两个子数组继续划分
		quickSort(arr, low, p-1)
		quickSort(arr, p+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// 获取一个枢纽值
	pivot := arr[high]
	// 先从左边找出第一个比p大的
	for low < high {
		for low < high && arr[low] <= pivot {
			low++
		}
		// 此时low比pivot大
		arr[high] = arr[low]
		// 继续移动high
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]
		// 继续交替
	}
	// 将枢轴值补充回去 此时应当low=high，所以哪个都可以
	arr[low] = pivot
	return low
}
