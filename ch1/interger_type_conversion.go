package main

import "fmt"

func integerTypeConversion() {
	var x int = 10
	var b byte = 100
	var sum3 int = x + int(b)
	var sum4 byte = byte(x) + b
	fmt.Println(sum3)
	fmt.Println(sum4)
}
