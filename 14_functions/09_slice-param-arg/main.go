package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average(data)
	fmt.Println(n)
	nice := []float64{44,4,12,2132,121,12312,12,1221213}
	nce := ce(nice)
	fmt.Sprint(nce)
}
func ce(sf []float64) float64 {
	to
}

func average(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
