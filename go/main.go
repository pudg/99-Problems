package main

import (
	"fmt"
	"pudg/99-problems/thirty"
	"pudg/99-problems/twenty"
)

func main() {
	fmt.Println(twenty.Pangram("We promptly judged antique ivory buckles for the next prize"))
	fmt.Println(twenty.ArrayPerm([]int{0, 1}, []int{0, 2}, 1))
	twenty.SalesMatch([]int{1, 2, 1, 2, 1, 3, 2})
	fmt.Println(twenty.BookPage(5, 3))

	fmt.Println(thirty.BFS(&thirty.Node{}, 22))
}
