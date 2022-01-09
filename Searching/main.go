package main

import (
	"fmt"

	bs "github.com/imrushi/algorithms/Searching/BinarySearch"
)

func main() {
	// arr := []int{-18, -12, -4, 0, 2, 3, 4, 15, 16, 18, 22, 45, 89}
	arr := []int{99, 80, 75, 22, 11, 10, 5, 2, -3}
	target := 22
	// ans := bs.BinarySearch(arr, target)

	ans := bs.OrderAgnosticBS(arr, target)

	fmt.Println(ans)
}
