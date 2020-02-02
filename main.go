package main

import "fmt"

func main() {
	a := []int{-10, 123, 34, 123, 432, 4, 23, 4, 3, 42, 342, 1, 23, 4, 23, 4123, 234}
	// a := []int{3, 4}
	quick_sort(a, 0, len(a))
	fmt.Println(a)
}
