# Go Programming Cheatsheet

## Basic Syntax

### Package Declaration
```go
package main  // Main package for executable programs
package math  // Package for library code
```

### Imports
```go
import "fmt"
import (
    "os"
    "strconv"
    "math/rand"
)
```

### Variables
```go
var name string = "John"
var age int = 25
city := "New York"  // Short declaration (inside functions only)
```

### Data Types
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `bool`
- `string`
- `rune` (Unicode code point)

### Functions
```go
func functionName(param1 type1, param2 type2) returnType {
    // function body
    return value
}

// Multiple return values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### Control Structures

#### If-Else
```go
if condition {
    // code
} else if anotherCondition {
    // code
} else {
    // code
}
```

#### For Loops
```go
// Traditional for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style loop
for condition {
    // code
}

// Range over slice/map
for index, value := range slice {
    fmt.Println(index, value)
}
```

#### Switch
```go
switch variable {
case value1:
    // code
case value2:
    // code
default:
    // code
}
```

### Arrays and Slices
```go
// Array (fixed size)
var arr [5]int
arr := [5]int{1, 2, 3, 4, 5}

// Slice (dynamic size)
var slice []int
slice := []int{1, 2, 3}
slice = append(slice, 4)

// Slicing
subSlice := slice[1:3]  // Elements at index 1 and 2
```

### Maps
```go
m := make(map[string]int)
m["key1"] = 10
m["key2"] = 20

value, exists := m["key1"]  // Check if key exists
```

### Structs
```go
type Person struct {
    Name string
    Age  int
    Address Address
}

type Address struct {
    Street string
    City   string
}

p := Person{Name: "John", Age: 25}
```

### Methods
```go
func (p Person) String() string {
    return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

func (p *Person) HaveBirthday() {
    p.Age++
}
```

### Interfaces
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

### Error Handling
```go
result, err := someFunction()
if err != nil {
    log.Fatal(err)
}
```

### Goroutines and Channels
```go
// Goroutine
go func() {
    fmt.Println("Running in background")
}()

// Channel
ch := make(chan string)
ch <- "message"
msg := <-ch

// Buffered channel
ch := make(chan int, 100)
```

### Pointers
```go
var x int = 42
p := &x  // Pointer to x
fmt.Println(*p)  // Dereference pointer

*p = 99  // Change value through pointer
```

## Useful Packages

- `fmt` - Formatted I/O
- `os` - Operating system interface
- `strconv` - String conversions
- `math` - Mathematical functions
- `time` - Time and date
- `strings` - String manipulation
- `sort` - Sorting algorithms
- `encoding/json` - JSON encoding/decoding
- `net/http` - HTTP client and server
- `database/sql` - SQL database interface

## Go Commands

```bash
go run main.go          # Run program
go build                # Build executable
go test                 # Run tests
go mod init module_name # Initialize module
go get package_name     # Add dependency
go fmt                  # Format code
go vet                  # Analyze code
```

## Conventions

- Use `gofmt` to format code
- Exported names start with capital letter
- Unexported names start with lowercase letter
- Use meaningful variable names
- Keep functions small and focused
- Handle errors explicitly
