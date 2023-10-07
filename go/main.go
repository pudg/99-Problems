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
	fmt.Printf("GroupedAnagrams: {%v}\n", ten.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Printf("TopKFreqs: {%v}\n", ten.TopKFreq([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, 2))
	fmt.Printf("RemoveElement: %v\n", ten.RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	fmt.Printf("ContainsDupsTwo: %v\n", ten.ContainsDupTwo([]int{1, 2, 3, 1}, 3))
	fmt.Printf("ContainsDupsTwo: %v\n", ten.ContainsDupTwo([]int{1, 2, 3, 1}, 3))
	fmt.Printf("ContainsDupsTwo: %v\n", ten.ContainsDupTwo([]int{1, 0, 1, 1}, 1))
	fmt.Printf("CombinationSum: %v\n", ten.CombinationSum([]int{2, 3, 5, 8}, 8))
	fmt.Printf("Cypher3: %v\n", ten.Cypher("abcdefghijklmnopqrstuvwxyz", 3))
	fmt.Printf("Cypher3: %v\n", ten.Cypher("There's-a-starman-waiting-in-the-SKY", 3))
	fmt.Printf("IsPalindrome: %v\n", ten.IsPalindrome(121))
}
