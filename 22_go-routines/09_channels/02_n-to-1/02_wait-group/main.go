package main

import (
	"fmt"
)

func send(out chan<- int) {
	out <- 89
	fmt.Println("写入",out)
	close(out)
}

func send1(out chan<- int){
	out <- 90
	close(out)
}
func recv1(in <-chan  int) {
	n := <-in
	fmt.Println(n)
}
func recv(in <-chan int) {
	n := <-in
	fmt.Println("读到", n)
}

func main() {
	ch := make(chan int)
	cn := make(chan int)
	go func() {
		send(ch)
	}()
	go func() {
		send(cn)
	}()
	recv(ch)
	recv(cn)
}