package main

import "fmt"

func main() {
	println("start main")
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// 如果不关闭channel,会引发panic
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
	println("end main")
}