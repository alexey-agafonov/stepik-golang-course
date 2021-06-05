package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/228263/step/3?unit=200796
func main() {
	var a, b, sum int

	fmt.Scan(&a, &b)

	for ; a <= b; a++ {
		sum += a
	}

	fmt.Println(sum)
}
