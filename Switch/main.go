package main

import "fmt"

func main() {
	i:=1
	switch i{
	case 1:
		fmt.Println("Value is 1")
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
	case 4:
		fmt.Println("Value is 4")
	default :
		fmt.Println("No value is found")
	}
}