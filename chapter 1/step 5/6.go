package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/228261/step/16?unit=200794
func main() {
	var d, h, m int
	
	fmt.Scan(&d)

	h = d / 30
	m = d % 30 * 2

	fmt.Println("It is", h, "hours", m, "minutes.")
}
