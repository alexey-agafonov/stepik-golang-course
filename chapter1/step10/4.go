package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/228263/step/5?unit=200796
func main() {
	var n, count, max int

	for fmt.Scan(&n); n != 0; fmt.Scan(&n) {
		switch {
		case n > max:
			max, count = n, 1
		case n == max: 
			count++
		}
	}

	fmt.Println(count)
}
