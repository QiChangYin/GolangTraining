package main

import "fmt"

func main() {
	i := 0
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
	k := 1
	for {
		fmt.Println(k)
		if k >= 100 {
			break
		}
		k++
	}
}
