package main

import "fmt"

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
