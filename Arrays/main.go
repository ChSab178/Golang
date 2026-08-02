package main

import "fmt"

// What is an array?
// Array is numbered sequence of specific length.
func main() {
	var num[5]int // declared an array.
	var str[5]string 
	var isBool[5]bool
	var Float[5]float32
	// What is stored when we will declared an array?
	fmt.Println(num) // Stores 0's
	fmt.Println(str) // Stores empty strings or spaces.
	fmt.Println(isBool) // Stores 'false'.
	fmt.Println(Float) // Stores 0's too
	
	// How to access a value of an array individually?
	// Syntax:
	// arrName[index number] index number always start from 0 to (array length - 1).
	fmt.Println(num[0])

	// How to change or add a value in an array individually?
	// Syntax:
	// arrName[index number]=value
	num[0]=2
	fmt.Println(num)
	

	// How to declare an array at the same time as initialization?
	arr:=[3]string{"a","b","c"}
	fmt.Println(arr)

	// How to create 2d arrays?
	arr2d:=[2][2]string{{"a","b"},{"c","d"}}
	fmt.Println(arr2d)

	// Benefits of Arrays:
	// - Fixed Size,that is predictable.
	// - Memory optimization.
	// - Constant time access.
}