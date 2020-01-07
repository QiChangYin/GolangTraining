package main

import "fmt"

func main() {

	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)

	nice := []int{1,2,3,4,5}
	nice1 := []int{5,6,7,8,8}
	nice = append(nice,nice1...)
	fmt.Println(nice)
}
