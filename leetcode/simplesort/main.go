package main

import "fmt"

// 有序数组插入元素，返回插入位置，
func main() {
	arr := []int{1, 2, 3, 5, 6}
	target := 4
	index := search2(arr, 0, len(arr)-1, target)

	fmt.Println(index)
	front := arr[:index]
	after := arr[index:]
	after2 := make([]int, len(after))
	copy(after2, after)
	fmt.Println(after)
	res := append(front, target)
	res = append(res, after2...)
	fmt.Println(res)
}

func search(num []int, start, end, target int) int {
	if len(num) == 0 {
		return 0
	}

	if start > end {
		return end + 1
	}

	// 划分
	mid := (end-start)/2 + start
	fmt.Println(mid)
	if num[mid] == target {
		return mid
	}
	if num[mid] < target {
		return search(num, mid+1, end, target)
	} else {
		return search(num, start, mid-1, target)
	}
	return -1
}

func search2(num []int, start, end, target int) int {
	for start <= end {
		mid := (end-start)/2 + start
		if num[mid] == target {
			return mid
		}
		if num[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
