package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	n := len(arr)
	res := [][]int{}
	visited := make([]bool, n)
	var dfs func(arr []int, n, index int, cur []int)
	dfs = func(arr []int, n, index int, cur []int) {
		if index >= len(arr) {
			return
		}
		// index 表示下一个
		// len 表示还要遍历几个
		// cur 表示当前组成的结果
		if n == 0 {
			res = append(res, cur)
			return
		}
		for i := range arr {
			if i == index || visited[i] {
				continue
			}
			visited[i] = true
			tmp := append(cur, arr[i])
			k := n - 1
			dfs(arr, k, i, tmp)
			visited[i] = false
		}
	}
	for i := range arr {
		visited[i] = true
		dfs(arr, n-1, i, []int{arr[i]})
		visited[i] = false
	}
	fmt.Println(res)
}
