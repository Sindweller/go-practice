package main

import "fmt"

func main() {
	header := map[string]string{"name": "1", "name2": "2"}
	resp := map[string]*string{}
	for k, v := range header {
		resp[k] = &v
	}
	fmt.Println(resp)
	// 打印输出：map[name:0xc00008e230 name2:0xc00008e230] 指向同一个地址，此时值全都会是最后一个遍历到的元素
	resp2 := map[string]string{}
	for k, v := range header {
		resp2[k] = v
	}
	fmt.Println(resp2)
	// 打印输出：map[name:1 name2:2] 是遍历并保存了所有的值
}
