package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	b := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	go func() {
		for i := 100; i < 10000 ; i++{

		}
	}

	for n := range c {
		fmt.Println(n)
	}
}
