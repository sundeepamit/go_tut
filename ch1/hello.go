package main //main package in go module contains a code that starts a program

import "fmt"

func main() {
	// calls Println() function from a fmt package.
	// fmt.Println("Hello World!!!")
	fmt.Printf("Hello, %s!!!!\n", "world")

	// Type conversion
	var x int64 = 10
	var y float64 = 30.12
	// converting int to float
	var sum1 float64 = float64(x) + y
	var sum2 int64 = x + int64(y)
	fmt.Println("Sum1 is ", sum1)
	fmt.Println("Sum2 is ", sum2)
	integerTypeConversion()
}
