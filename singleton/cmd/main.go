package main

import (
	"fmt"
	"patterns/singleton"
)

func main() {
	// Get the singleton instance
	s1 := singleton.GetInstance()
	fmt.Println(s1.Data)

	// Get the same instance again
	s2 := singleton.GetInstance()
	fmt.Println(s2.Data)

	// Verify it's the same instance
	fmt.Println(s1 == s2) // Output: true

}
