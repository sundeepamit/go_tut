package main

import "fmt"

func varribaleDeclaration() {
	// Declaring a variable using var keyword.
	// this is most verbose way to declare a variable.
	var age int = 29
	var x = 10 //type infrence
	fmt.Println(age, x)
	// declare a multiple variable at once with var keyword
	var name, city string = "amit", "Tyoko"
	fmt.Println(name, city)
	// you can declare all zero values of same type
	var a, b int
	fmt.Println(a, b)
	x = 23
	// go also support a short declaration and  assignment operator
	// := this only work inside function.
	hobby := "coding" //same a var hobby = "coding"
	fmt.Println(hobby)
	x, y := 10, "hello"
	fmt.Println(x, y)
}
