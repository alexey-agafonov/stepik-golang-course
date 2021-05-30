package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/232593/step/6?thread=solutions&unit=205068
func main() {
	var b int

	fmt.Scan(&b)

	f := b / 100
	s := b / 10 % 10
	t := b % 10

	if f == s || f == t || s == t {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
