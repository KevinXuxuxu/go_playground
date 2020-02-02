package main

import (
	"fmt"
)

func main() {
	a := []int{6, 3, 4, 2, 5, 7, 10000, 2, 3422, 123, 12341234, 234, 22, 341234}
	for i := len(a) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	fmt.Println(a)
}
