package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/232593/step/5?unit=205068
func main() {
    var c int
    
	fmt.Scan(&c)

    switch {
        case c > 0: fmt.Println("Число положительное") 
        case c < 0: fmt.Println("Число отрицательное")
        default: fmt.Println("Ноль")
    }
}
