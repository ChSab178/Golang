package main

import (
	"fmt"
	"maps"
)

// Maps are like the dictionaries in Python and objects like in JavaScript.
func main() {
	m:=make(map[string]string)
	m["name"]="golang"
	m["grade"]="A"
	// Get an element:
	fmt.Println(m["name"])
	fmt.Println(m)
	// If you try to get the element that is not in the map it will return the empty string or 0 or false based on the dtype of map.
	a:=make(map[string]int)
	a["Age"]=17
	a["Class"]=11
	fmt.Println(a)

	// How to get the length of the map:
	fmt.Println(len(a))

	// How to delete an element in the Maps:
	delete(a,"Class")
	fmt.Println(a)

	// How to remove all elements in the map:
	clear(a)
	fmt.Println(a)

	// How to initialize and and declare the map in one line:
	b:=map[string]int{"Age":17,"Class":11}
	fmt.Println(b)

	// How to compare a map:
	// As we used slices in the slice to check does the first slice is equal to the other one. Same as we used package called maps.
	fmt.Println(maps.Equal(a,b))

	// Range function with Maps:
	c:=map[string]string{"name":"Hammad","fname":"Asghar","Class":"Intermediate"}
	for key,value:=range c{
		fmt.Println(key," : ",value)
	}

	// Range function with Strings:
	e:="Hammad Asghar"
	for i,value:=range e{
		fmt.Println(i+1," : ",string(value)) // It will returns the ASCII values that's why we convert the ASCII values to the string characters
	}
}