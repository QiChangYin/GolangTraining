package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i >= 50 {
			break
		}
	}

	j := 1
	for {
		j++
		if j % 2 == 0 {
			fmt.Println(j)
			continue
		}
		fmt.Println(j)
		if j >= 100 {
			break
		}
	}
}
