package main

import "fmt"

var e int = 6      // The normal method of creating var allows us to write variable outside the function.
const age int = 30 // constant can also be created outside the function.
func main() {
	var a int = 1 // It stores int value.
	var b = 2     // It automatically assigns the type to the var based on the value.
	c := 3        // ShortHand method to write the var and giving it's value.
	b = 4         // You can change the value.
	d := 5        // ShortHand only can be declared inside the function.
	const val=40  // Const are used when we don't want to change the value of var.
	// val=50 Value of constant cannot be changed.

	// Here's how we declare multiple const vars at the same time professionally:
	const (
		port=3000
		host="localhost"
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(port)
	fmt.Println(host)
}
