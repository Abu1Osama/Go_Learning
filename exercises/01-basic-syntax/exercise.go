package main

import "fmt"

// Exercise: Complete the following functions to practice Go basics

// TODO: Fix this function to properly greet someone
func greet(name string) {
	// Your code here - should print "Hello, [name]!"
}

// TODO: Create a function that calculates the area of a circle
// Formula: area = π * r²
func circleArea(radius float64) float64 {
	// Your code here
}

// TODO: Create a function that checks if a number is even
func isEven(num int) bool {
	// Your code here
}

// TODO: Create a function that finds the maximum of three numbers
func maxOfThree(a, b, c int) int {
	// Your code here
}

func main() {
	fmt.Println("=== Go Exercises ===")
	
	// Test your functions here
	greet("Go Developer")
	
	area := circleArea(5.0)
	fmt.Printf("Area of circle with radius 5: %.2f\n", area)
	
	fmt.Printf("Is 4 even? %t\n", isEven(4))
	fmt.Printf("Is 7 even? %t\n", isEven(7))
	
	fmt.Printf("Max of (3, 7, 5): %d\n", maxOfThree(3, 7, 5))
	fmt.Printf("Max of (10, 2, 8): %d\n", maxOfThree(10, 2, 8))
}
