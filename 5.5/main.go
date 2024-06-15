package main

import (
	"fmt"
)

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	numbers := make([][]string, n)
	for i := 0; i < len(numbers); i++ {
		numbers[i] = make([]string, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_, _ = fmt.Scan(&numbers[i][j])
		}
	}
	isBnW := true
outro:
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if numbers[i][j] == "W" || numbers[i][j] == "B" || numbers[i][j] == "G" {
				continue
			} else {
				isBnW = false
				break outro
			}
		}
	}
	if isBnW == true {
		fmt.Println("#Black&White")
	} else {
		fmt.Println("#Color")
	}
}
