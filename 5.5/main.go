package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	numbers := make([]int, n)
	counter := 1
	for i := 0; i < n; i++ {
		_, _ = fmt.Scan(&numbers[i])
	}

	for j := 0; j < n-1; j++ {
		if numbers[j] < numbers[j+1] {
			counter++
		}
	}
	fmt.Println(counter)
}
