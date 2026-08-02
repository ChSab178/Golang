package main

import "fmt"

// There is only for loop in go only.
func main() {
	// We used the for loop to create while loop as:
	n:=1
	for n<=3{
		fmt.Println(n)
		n++
	}

	// Infinite loop:
	// for{
	// 	fmt.Println("This is infinite loop....") 
	// }

	// Classic for loop:
	for i:=0;i<3;i++{
		fmt.Println(i)
	}

	// For loop with range:
	for l:=range 10{
		fmt.Println(l)
	}
}

