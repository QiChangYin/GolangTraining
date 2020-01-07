package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(myGreeting["Jenny"])
	myGreeting1 := map[string]string{
		"Timw": "12121",
		"asadas":"saddsaads",
    }
    fmt.Println(myGreeting1["Timw"])
}
