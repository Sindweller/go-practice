package main

import "fmt"

func main() {
	board := make([][]byte, 3)
	board[0] = []byte{'A', 'B', 'C', 'E'}
	board[1] = []byte{'S', 'F', 'C', 'S'}
	board[2] = []byte{'A', 'D', 'E', 'E'}
	fmt.Println(board)
	fmt.Println(exist(board, "ABCB"))
}

func exist(board [][]byte, word string) bool {
	// visited 记录是否已访问，因为每个cell只能用一次
	// 直接回溯 方向可以是上下左右，超过边界不能访问，visited也不能访问
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}

	var dfs func([][]byte, int, int, string) bool
	dfs = func(board [][]byte, i, j int, subword string) bool {
		fmt.Println(i, j)
		fmt.Println(subword)
		// 边界条件 这里不能或，需要立刻返回
		if len(subword) == 0 {
			return true
		}
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || visited[i][j] {
			fmt.Println("超出边界")
			return false
		}
		if len(subword) == 1 && board[i][j] == subword[0] {
			return true
		}

		// 如果是错误的不需要递归
		if board[i][j] != subword[0] || visited[i][j] {
			return false
		}
		// 继续
		fmt.Println("未完成...")
		if board[i][j] == subword[0] {
			fmt.Println("ok")
			//上
			visited[i][j] = true
			if dfs(board, i-1, j, subword[1:]) || dfs(board, i+1, j, subword[1:]) || dfs(board, i, j-1, subword[1:]) || dfs(board, i, j+1, subword[1:]) {
				return true
			}
		}
		fmt.Println("false")
		visited[i][j] = false
		return false
	}
	for i := range board {
		for j := range board[0] {
			if dfs(board, i, j, word) {
				return true
			}
		}
	}
	return false
}
