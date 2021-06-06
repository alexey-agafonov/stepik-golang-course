package main

import "fmt"

func YearCount(x int, p int, y int) int {
	var year int

	for x <= y {
		x = x + x * p / 100
		year++
	}

	return year
}

func main() {
	var p, x, y int
	fmt.Scan(&x, &p, &y)
	fmt.Println(YearCount(x, p, y))
}
