package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/228263/step/4?unit=200796
func main() {
	var n, sum int

	for fmt.Scan(&n); n > 0; n-- {
		var number int

		if fmt.Scan(&number); number >= 10 && number <= 99 && number % 8 == 0 {
			sum += number
		}
	}

	fmt.Println(sum)
}
