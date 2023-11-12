package main

import (
	"fmt"
	"pudg/99-problems/twenty"
)

func main() {
	fmt.Println(twenty.PlusMinus([]int{-3, -2, -1, 0, 1, 2, 3}))
	fmt.Println(twenty.MinMaxSum([]int{7, 9, 3, 5, 1}))
	fmt.Println(twenty.TimeConvert("02:01:00PM"))
	fmt.Println(twenty.TimeConvert("12:01:00PM"))
	fmt.Println(twenty.TimeConvert("12:01:00AM"))
}
