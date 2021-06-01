package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/228263/step/8?thread=solutions&unit=200796
func main() {
	var n int

	for fmt.Scan(&n); n <= 100; fmt.Scan(&n) {
		if n > 10 {
			fmt.Println(n)
		}
	}
}
