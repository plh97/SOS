package main

import (
	"fmt"
)

func fib(n int, c chan float32) {
	x, y := 0.0, 1.0
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func main() {
	c := make(chan float32, 10009988.0)
	go fib(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
