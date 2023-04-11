package main

import "fmt"

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	return search(nums)
}

func search(nums []int) [][]int {
	var res [][]int
	for i := range nums {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		if l >= r {
			break
		}
		// 跳过
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if s == 0 {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				fmt.Println(res)
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			} else if s < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot := arr[left]
	i, j := left, right
	for i != j {
		//j 大 i 小
		for i < j && arr[j] >= pivot {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
		for i < j && arr[i] <= pivot {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	quickSort(arr, left, i-1)
	quickSort(arr, i+1, right)
}
