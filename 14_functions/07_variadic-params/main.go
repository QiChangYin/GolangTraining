package main

import "fmt"

func main() {
	//n := average(43, 56, 87, 12, 45, 57)
	//fmt.Println(n)

	c := nice(43,3,4,3,2,12,2,12,12,12,21,21,123,12)
	fmt.Sprint(c)
}

func nice(cd ...float64) float64{
	print(cd)
	var to float64
	for _, v := range cd {
		to += v
	}
	return to / float64(len(cd))
}
//
//func average(sf ...float64) float64 {
//	fmt.Println(sf)
//	fmt.Printf("%T \n", sf)
//	var total float64
//	for _, v := range sf {
//		total += v
//	}
//	return total / float64(len(sf))
//}

