package main

import "fmt"

func main() {
	stack := getArrayStack(10)
	fmt.Println(stack)
	for i := 0; i < 10; i++ {
		stack.push(i)
		fmt.Println(stack)
	}

	stack.pop()
	fmt.Println(stack)
}
