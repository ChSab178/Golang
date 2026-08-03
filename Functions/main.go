package main

import "fmt"

// If we have same dtype we can specify it at once:
// func add(a,b int){...} LIke that.

func add(a int,b int) int{
	return a+b 
}

func getLanguages()(string,string,string,string){
	return "Python","Cpp","JavaScript","Golang"
}

// We can also pass the functions as arguments in the functions:
func Square(fn func(a int) int,b int)int{
	return fn(4) * 2
}

func main() {
	result:=add(2,3)
	fmt.Println(result)

	lang1,lang2,_,lang4:=getLanguages() // If you don't want to use a var just put _ instead of it's name.
	fmt.Println(lang1,lang2,lang4)
	// You can directly print the values:
	fmt.Println(getLanguages())

	// You can create the function in the variable:
	fn:=func (a int)int{
		return a
	}
	value:=Square(fn,2)
	fmt.Println(value)
}