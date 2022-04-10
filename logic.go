package main

import (
	"fmt"
)

func main() {
	var length int
	fmt.Print("Total Numbers: ")
	fmt.Scan(&length)
	numbers := make([]int, length)
	for a := 0; a < length; a++ {
		fmt.Print("Numbers-", a, " : ")
		fmt.Scan(&numbers[a])
	}
	var target int
	fmt.Print("Target: ")
	fmt.Scan(&target)

	fmt.Println(numbers)
	fmt.Println(target)

	numMap := map[int]int{}

	for a := 0; a < len(numbers); a++ {
		complement := target - numbers[a]
		if j, ok := numMap[complement]; ok {
			fmt.Print(j, ",", a)
		}
		numMap[numbers[a]] = a
	}
}