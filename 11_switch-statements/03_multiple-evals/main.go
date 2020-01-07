package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny", "dsdsa":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Marcus", "Medhi", "dadasfsad":
		fmt.Println("Both of your names start with M")
	case "Julian", "Sushant","saddsafasd":
		fmt.Println("Wassup Julian / Sushant")
	}
}
