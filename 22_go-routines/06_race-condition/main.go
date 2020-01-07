package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func main() {
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	go incrementor("cni")
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}


func incrementor(s string) {
	rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// go run -race main.go
// vs
// go run main.go
