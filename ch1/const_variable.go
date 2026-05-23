package main

import "fmt"

/*
*
Contants are simply a way to attach a meaningful name to raw value(literals)

	A literal is a raw hardcoded value written directly in code.
	for eg:
	3.14
	"hello"
	100
	true

	A Constant give that literal a name
	const PI = 3.14  //"PI" is the name ,3.14 is a literal
	const greeting ="hello"

	So const doesn't do any computation or logic — it purely gives a human-readable label to a value that would otherwise just be a "magic number/string" sitting in your code.
*/

// Constants in Go are a way to give names to literals. There is no way
// in Go to declare that a variable is immutable.

// const gives you immutability only for simple literal values. For variables holding structs, slices, maps etc. — Go gives you no immutability mechanism at all.
const pi float64 = 3.14
const (
	idkey   = "id"
	nameKey = "name"
)
const z = 20 * 10

func constVariables() {
	const y = "hello"
	fmt.Println(pi)
	fmt.Println(y)
	// pi = 23.1  this will not compile
	// y = "world"

	// fmt.Println(pi)
	// fmt.Println(y)
}
