package main

import "fmt"

// Basic function with parameters and return value
func add(a int, b int) int {
	return a + b
}

// Function with multiple return values
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return dividend / divisor, nil
}

// Function with named return values
func calculateRectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return - returns area and perimeter
}

// Variadic function (accepts variable number of arguments)
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Recursive function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Function as a value (higher-order function)
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	fmt.Println("=== Functions Demo ===")
	
	// Basic function call
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)
	
	// Multiple return values
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %.1f\n", quotient)
	}
	
	quotient, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	
	// Named return values
	area, perimeter := calculateRectangle(5, 3)
	fmt.Printf("Rectangle (5x3): Area=%.1f, Perimeter=%.1f\n", area, perimeter)
	
	// Variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Printf("Sum of 1,2,3,4,5 = %d\n", total)
	
	// Recursive function
	fact := factorial(5)
	fmt.Printf("5! = %d\n", fact)
	
	// Higher-order function
	sumFunc := func(a, b int) int { return a + b }
	multiplyFunc := func(a, b int) int { return a * b }
	
	fmt.Printf("Apply sum to 4,6: %d\n", applyOperation(4, 6, sumFunc))
	fmt.Printf("Apply multiply to 4,6: %d\n", applyOperation(4, 6, multiplyFunc))
}
