package main

import "fmt"

type MinStack struct {
	stack []int64
	min   int64
}

/** initialize your data structure here. */
func Constructor() MinStack {
	maxInt := int64(int(^uint(0) >> 1))
	return MinStack{
		min: maxInt,
	}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
	} else {
		tmp := int64(x) - this.min
		this.stack = append(this.stack, tmp)

	}
	if int64(x) < this.min {
		this.min = int64(x)
	}
	fmt.Println("push")
	fmt.Println(this.stack)
	fmt.Println(this.min)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	if this.stack[len(this.stack)-1] < 0 {
		this.min = this.min - this.stack[len(this.stack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int64 {
	if len(this.stack) == 0 {
		return 0
	}
	fmt.Println("top")
	fmt.Println(this.stack[len(this.stack)-1])
	fmt.Println(this.min)
	fmt.Println(this.min + this.stack[len(this.stack)-1])
	if this.stack[len(this.stack)-1] < 0 {
		return this.min
	} else {
		return this.min + this.stack[len(this.stack)-1]
	}
}

func (this *MinStack) Min() int64 {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
