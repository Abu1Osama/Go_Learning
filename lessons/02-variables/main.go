package main

import "fmt"

func main() {
	// Variable declarations in Go
	
	// Method 1: Using var keyword with type
	var name string = "John Doe"
	var age int = 25
	var height float64 = 5.9
	var isStudent bool = true
	
	// Method 2: Type inference (Go determines the type)
	var city = "New York"
	var score = 95.5
	
	// Method 3: Short declaration (only inside functions)
	country := "USA"
	grade := 'A' // rune type (Unicode code point)
	
	// Multiple variable declaration
	var (
		x int = 10
		y int = 20
		z int = 30
	)
	
	// Zero values (variables declared without initialization)
	var defaultInt int
	var defaultString string
	var defaultBool bool
	
	// Print all variables
	fmt.Println("=== Variables Demo ===")
	fmt.Printf("Name: %s, Age: %d, Height: %.1f, Student: %t\n", name, age, height, isStudent)
	fmt.Printf("City: %s, Score: %.1f\n", city, score)
	fmt.Printf("Country: %s, Grade: %c\n", country, grade)
	fmt.Printf("Coordinates: (%d, %d, %d)\n", x, y, z)
	fmt.Printf("Zero values: int=%d, string='%s', bool=%t\n", defaultInt, defaultString, defaultBool)
	
	// Constants
	const PI = 3.14159
	const GREETING = "Hello"
	
	fmt.Printf("Constants: PI=%.5f, GREETING=%s\n", PI, GREETING)
}
