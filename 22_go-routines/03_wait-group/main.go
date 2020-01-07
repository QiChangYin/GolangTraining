package main

import (
	"fmt"
	"sync"
)
// 创建同步等待组对象
var wg sync.WaitGroup
//WaitGroup: 同步等待阻塞
//可以使用Add()， 设置等待组中执行的子goroutine的数量
//在main函数中，使用wait(),让主程序处于等待状态
// 直到等待组中的程序执行完毕，解除阻塞
// 子goroutine对应的函数中，wg.Done(),用于让等待组中的子程序数量减1，
// 设置等待组中，要执行的goroutine的数量

func main() {
	wg.Add(3)
	go foo()
	go bar()
	go cao()
	wg.Wait()
}


func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
	wg.Done()
}

func cao() {
	for j := 0 ; j < 45; j++ {
		fmt.Println("CAO",j)
	}
	wg.Done()

	// 给wg等待中执行的goroutine的数量减1，通ADD(-1)
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
