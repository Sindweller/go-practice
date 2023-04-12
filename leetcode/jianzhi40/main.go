package main

import "fmt"

func main() {
	nums := []int{3, 2, 1}
	k := 2
	heap := []int{}
	for i := range nums {
		if len(heap) < k {
			heap = append(heap, nums[i])
		} else {
			if nums[i] < heap[0] {
				heap[0] = nums[i]
				heapify(heap, k, 0)
			}
		}
	}
	fmt.Println(heap)
}

func heapSort(arr []int) []int {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	//for i := n - 1; i >= 0; i-- {
	//	arr[0], arr[i] = arr[i], arr[0]
	//	heapify(arr, i, 0)
	//}
	return arr
}

func heapify(arr []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := l + 1
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, i)
	}
	return
}
