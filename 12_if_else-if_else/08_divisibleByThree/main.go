package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	for j := 1 ; j <= 1000 ; j++ {
		if j % 3 == 0 {
			fmt.Println(j)
		}
	}
}
