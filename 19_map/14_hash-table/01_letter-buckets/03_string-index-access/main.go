package main

import "fmt"

func main() {
	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)

	nice := "caonice"
	fmt.Println(rune(nice[0]))
}
