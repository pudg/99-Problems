package main

import (
	"fmt"
	"pudg/99-problems/ten"
)

func main() {
	fmt.Printf("RemoveDups: k = {%v}\n", ten.RemoveDups([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Printf("TwoSum: {%v}\n", ten.TwoSumV2([]int{3, 2, 4}, 6))
	fmt.Printf("SquareRoot of {%v} = {%v}\n", 327, ten.NewtonSquareRoot(327, 0.00001))
	fmt.Printf("ContainsDups: {%v}\n", ten.ContainsDup([]int{1, 2, 3, 4}))
	fmt.Printf("ContainsDups: {%v}\n", ten.ContainsDup([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
