package main

import (
	"fmt"
	"bytes"
)

func removeElements(head *ListNode, val int, depth int) *ListNode {
	// depthString 代表递归深度
	depthString := generateDepthString(depth)
	// 递归调用前打印
	fmt.Print(depthString)
	fmt.Println("Call: remove", val, " in ", head)

	if head == nil {
		fmt.Print(depthString)
		fmt.Println("Return: ", head)
		return nil
	}

	res := removeElements(head.Next, val, depth + 1)
	// 递归调用后打印
	fmt.Print(depthString)
	fmt.Println("After remove ", val, ": ", res)

	ret := &ListNode{}
	if head.Val == val {
		ret = res
	} else {
		head.Next = res
		ret = head
	}
	// 节点处理后打印
	fmt.Print(depthString)
	fmt.Println("Return: ", ret)

	return ret
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}

func main()  {
	nums := []int{1, 2, 6, 3, 4, 5, 6}
	head := GetListNode(nums)
	fmt.Println(head)

	res := removeElements(head, 6, 0)
	fmt.Println(res)
}