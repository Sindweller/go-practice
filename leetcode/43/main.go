package main

import "fmt"

func main() {
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	binarySearch([]int{1, 2}, 2, false)
}

func searchRange(nums []int, target int) []int {
	t := binary(0, len(nums)-1, target, nums)
	if t < 0 {
		return []int{-1, -1}
	}
	i := t
	for i >= 0 {
		fmt.Println(i)

		if nums[t] == nums[i] {
			i--
		} else {
			fmt.Println(i)
			fmt.Println("is not")
			i++
			break
		}
	}

	var res []int
	res = append(res, i)
	fmt.Println(res)
	j := t
	for j < len(nums) {
		fmt.Println(j)

		if nums[t] == nums[j] {
			j++
		} else {
			fmt.Println(j)
			fmt.Println("is not")
			j--
			break
		}
	}

	res = append(res, j)
	return res
}

func binary(left, right, target int, nums []int) int {
	fmt.Println(left, right, target)
	if left >= right {
		if left < len(nums) {
			return left
		} else {
			return -1
		}
	}

	mid := (left + right) / 2
	fmt.Printf("mid is %d\n", mid)
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binary(left, mid-1, target, nums)
	} else {
		return binary(mid+1, right, target, nums)
	}
	return -1
}

func binarySearch(nums []int, target int, lower bool) int {
	res := []int(nil)
	fmt.Println(res)
	for i := range res {
		fmt.Println(res[i])
		fmt.Println(&res[i])
	}
	fmt.Println(&res)
	fmt.Println(res == nil)
	return 1
}
