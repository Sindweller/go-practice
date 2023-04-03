package main

import "fmt"

// 目前这是一个小根堆
func heapify(arr []int, n, i int) {
	// n是目前堆的大小，必须传 因为并不是整个数组都需要被排序
	// 取当前元素
	largest := i
	// 取左右孩子
	l := 2*i + 1
	r := 2*i + 2

	// 取左右孩子中更大的那个
	if l < n && arr[l] > arr[largest] {
		largest = l
	}

	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	// 如果largest换位了
	if largest != i {
		// 进行调整
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest) // 继续往上调整
	}
}

// 外层的实现
func heapSort(arr []int) []int {
	n := len(arr)
	// 初始化大根堆，此时所有元素都无序，但是已经放在数组里了。
	// 注意调整的时候要比较父和子节点，所以我们应该从最后一个非叶子结点（有孩子的节点）开始调整，不然没有意义
	// 最后一个非叶子结点位于：n/2-1
	// 为什么是这个呢？因为假设一个节点i的父节点是(i-1)/2,所以如果这个节点是数组末尾的话，就是n-1,自然他的父节点就是(n-1-1)/2.也就是n/2-1
	for i := n/2 - 1; i >= 0; i-- {
		// 不断向上调整
		heapify(arr, n, i) // 这里只需要变动n以前的
	}
	// 上面的第一个for循环能保证把最大值调整到顶
	// 但是，上面从下往上调整之后，下面的子树又不一定满足大根堆性质了，所以还得从上往下调整一遍
	// 此时，直接把顶元素（最大的）与最后一个叶子节点交换，然后，直接对根节点（原来的最后一个节点）开始向下递归调整 以后每次都是这么做的
	for i := n - 1; i >= 0; i-- {
		// 将堆顶元素与最后的交换，然后再调整堆
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0) // 从顶向下调整
	}
	return arr
}

func main() {
	arr := []int{12, 11, 3, 2, 5, 7, 77, 6}
	fmt.Println("unsorted arr: ", arr)

	arr = heapSort(arr)

	fmt.Println("sorted array: ", arr)
	//unsorted arr:  [12 11 3 2 5 7 77 6]
	//sorted array:  [2 3 5 6 7 11 12 77]
	//这个是利用大根堆达到从小到大排序（每次获得最大的元素，然后放到最后）
}
