package main

import "fmt"

// The solution for an assignment
// https://stepik.org/lesson/232593/step/8?thread=solutions&unit=205068
func main() {
	var n, firstSum, lastSum int
	
	fmt.Scan(&n)

	firstSum = (n / 100000) + (n / 10000 % 10) + (n / 1000 % 10)
	lastSum = (n / 100 % 10) + (n / 10 % 10) + (n % 10)

	if firstSum == lastSum {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
