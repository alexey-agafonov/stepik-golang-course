package main

import "fmt"

func TripleText() string {
	s := "I like Go!\n"
	return s + s + s
}

// The solution for an assignment
// https://stepik.org/lesson/228260/step/3?unit=200793
func main() {
	fmt.Println(TripleText())
}
