package main

import "fmt"

func main() {

	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	canice := []string{
		"AAAA",
		"CCCC",
		"DDDD",
	}
	for i,ca := range canice {
		fmt.Println(i,ca)
	}
	for k := 0 ; k < len(canice); k++{
		fmt.Println(canice[k])
	}


	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}

}
