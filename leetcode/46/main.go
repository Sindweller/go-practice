package main

import "fmt"

func main() {
	//fmt.Println(contains([]int{6, 3, 1}, 6))
	fmt.Println(permute([]int{5, 4, 6, 2}))
}

func permute(nums []int) [][]int {
	res := [][]int{}
	var dfs func(cur []int)
	dfs = func(cur []int) {
		// 因为index是一个一个加的，所以判断==就返回，不会抵达>
		// if index > len(nums){
		//     return
		// }
		if len(cur) == len(nums) {
			res = append(res, append([]int(nil), cur...))
			return
		}

		// 对于每一个可能的选项
		for i := 0; i < len(nums); i++ {
			// 因为题目给出不含重复数字，所以可以用这种方法来简单判断是否可以选择
			if contains(cur, nums[i]) {
				continue
			}
			// 计入当前结果
			cur = append(cur, nums[i])
			fmt.Println(cur)
			// 继续下一个
			dfs(cur)
			// 减去当前结果，遍历到下一个分支
			cur = cur[:len(cur)-1]
		}

	}
	dfs([]int(nil))
	return res
}
func contains(a []int, target int) bool {
	for i := range a {
		if a[i] == target {
			return true
		}
	}
	return false
}

// 二分查找必须有序
func binarySearch(left, right, target int, nums []int) int {
	fmt.Println(left, right)
	if left > right {
		return -1
	}
	// if nums[left] == target{
	//     return left
	// }
	mid := (left + right) >> 1
	fmt.Printf("mid is %d\n", mid)
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return binarySearch(left, mid-1, target, nums)
	} else {
		return binarySearch(mid+1, right, target, nums)
	}
	return -1
}
