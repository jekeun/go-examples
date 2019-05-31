package main

import "fmt"

func main() {
	printSlice()
}

func printSlice() {
	testStr := []string{"A", "B", "C"}

	fmt.Printf("%v", testStr)
}
