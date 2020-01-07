package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")
	greeting = append(greeting, "Zǎo'ān")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))

	greetingNice := make([]string, 3,100)
	greetingNice[0] = "12"
	greetingNice[1] = "2"
	greetingNice[2] = "44"
	greetingNice = append(greetingNice,"1121212112")
	fmt.Println(greetingNice[2])
	fmt.Println(len(greetingNice))
	fmt.Println(cap(greetingNice))

}
