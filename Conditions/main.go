package main

import "fmt"

func main() {
  // With simple if condition:
  // Simple if is used when we have only 1 condition and we have to choose that condition or not as:
  grade:="A"
  if grade=="A"{
	fmt.Println("You are an incredible student")
  }	

  // With if-else:
  // if-else structure is used when we have 2 conditions and we have to choose 1 condition from it as :
  marks:=50
  if marks>=60{
    fmt.Println("Passed")
  }else{
    fmt.Println("Failed")
  }

  // With if-elseif-else:
  // if-elseif-else structure is used when we have more than 2 conditions and we have to choose 1 from it as :
  age:=71
  if age>=18 && age<=70{
    fmt.Println("You can drive sir.")
  }else if age>70 {
    fmt.Println("You are way too old why still want to drive.")
  }else{
    fmt.Println("You cannot drive Mr/Mr's teenager.")
  }
}