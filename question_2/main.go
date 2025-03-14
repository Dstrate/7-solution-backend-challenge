package main

import "fmt"

func main() {
	var input string
	fmt.Print("Enter Code (L,R,=) :")
	fmt.Scanln(&input)
	ans2 := DecodeLeftRightEqual(input)
	fmt.Printf("Answer for Question 2 is : %v\n", ans2)
}
