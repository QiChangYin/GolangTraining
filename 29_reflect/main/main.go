package main

import "fmt"

type Order struct{
	orderId int
	customer int
	callback func()
}

func main(){
	var i interface{}
	i = Order{
		orderId:  0,
		customer: 0,
		callback: nil,
	}
	value, ok := i.(Order)
	if !ok {
		fmt.Println("-----------")
	    return
	}
	fmt.Println("asdsadsa",value)
}