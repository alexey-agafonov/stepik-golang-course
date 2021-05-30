package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/228261/step/11?unit=200794
func main() {
	var a int

	fmt.Scan(&a)
	a = a * 2 + 100

	fmt.Print(a)
}
