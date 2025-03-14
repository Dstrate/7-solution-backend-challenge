package main

import "fmt"

func main() {
	filePath := "./src/hard.json"
	ans := FindMostValuePath(filePath)
	fmt.Printf("Answer for Question 1 is : %v\n", ans)
}
