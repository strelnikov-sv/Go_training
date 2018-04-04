package main

import "fmt"

func fibonacci() func() int {

	pre := 0
	cur := 1
	return func() int {
		pre, cur = cur, pre+cur
		return cur
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
