package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/232593/step/7?thread=solutions&unit=205068
func main() {
	var n int
	
	fmt.Scan(&n)
	
	switch {
	case n >= 10000: fmt.Println(n / 10000)
	case n >= 1000: fmt.Println(n / 1000)
	case n >= 100: fmt.Println(n / 100)
	case n >= 10: fmt.Println(n / 10)
	default: fmt.Println(n)
	}
}
