package main

import (
	"fmt"
)

func SelectionSort() []int {
	var currentMin int

	arr := []int{0, 2, 8, 5, 3, 9, 4, 1, 20, 10, 7, -1, 10}

	for i := 0; i < len(arr); i++ {
		j := i + 1
		for ; j < len(arr); j++ {
			currentMin = arr[i]
			if currentMin > arr[j] {
				currentMin = arr[j]
				arr[i], arr[j] = arr[j], arr[i]
				continue
			} else if currentMin < arr[j] {
				continue
			}
			break
		}
	}
	return arr
}

func suAkarYolunuBulurBeeFaktoriyel(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * suAkarYolunuBulurBeeFaktoriyel(n-1)
	}

}

func main() {
	fmt.Println(SelectionSort())
	fmt.Println(suAkarYolunuBulurBeeFaktoriyel(5))
}
