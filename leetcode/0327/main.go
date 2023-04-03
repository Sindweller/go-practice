package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

type LRU struct {
	//双向链表
	list []*Node
	// 哈希表
	mapp     map[int]*Node
	capacity int // 容量 超出就删除表尾
}

// 查询
func (l *LRU) Get(key int) (int, error) {
	//get的时候如果存在就放在链表尾
	p, ok := l.mapp[key]
	if !ok {
		return 0, errors.New("未找到")
	}
	//找到指针
	prev := p.Prev
	after := p.Next
	tail := l.list[len(l.list)-1]
	if prev != nil {
		prev.Next = p.Next
	}
	if after != nil {
		after.Prev = p.Prev
	}
	tail.Next = p
	l.list = l.list[1:]
	l.list = append(l.list, p)
	p.Next = nil
	p.Prev = tail
	return p.Val, nil
}

func (l *LRU) Add(val int) {
	if len(l.list) == l.capacity {
		// 删除
		l.DelHead()
	}

	cur := &Node{
		Val:  val,
		Next: nil,
	}
	if len(l.list) > 0 {
		cur.Prev = l.list[len(l.list)-1]
	}
	l.list = append(l.list, cur)
	l.mapp[val] = l.list[len(l.list)-1]
}

func (l *LRU) DelHead() {
	cur := l.list[1]
	cur.Prev = nil
	l.list[0].Next = nil
	remove := l.list[0].Val
	l.list = l.list[1:]
	delete(l.mapp, remove)
}

func main() {
	lru := &LRU{
		list:     make([]*Node, 0),
		mapp:     make(map[int]*Node),
		capacity: 3,
	}
	lru.Add(1)
	fmt.Println(lru.list)
	fmt.Println(*lru.list[0])
	lru.Add(2)
	fmt.Println(lru.list)
	fmt.Println(*lru.list[0], *lru.list[1])
	lru.Add(3)
	fmt.Println(lru.list)
	fmt.Println(*lru.list[0], *lru.list[1], *lru.list[2])
	lru.Get(1)
	fmt.Println(lru.list)
	fmt.Println(*lru.list[0], *lru.list[1], *lru.list[2])
	lru.Add(4)
	fmt.Println(lru.list)
	fmt.Println(*lru.list[0], *lru.list[1], *lru.list[2])
}
