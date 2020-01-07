package main

import "fmt"

func main() {
	c := incrementor()
	for n := range puller(c) {
		fmt.Println(n)
	}
	d := incrementor()
	for  l := range puller(d) {
		fmt.Println(l)
	}
}


func incrementor() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func new() <-chan int {
	out := make(chan int)
	go func() {
		for j := 0 ; j < 100 ; j++ {
			out <- j
		}
	}()
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
