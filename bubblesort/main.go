package main

import "fmt"

func main() {

	a := []int{2, 6, 1, 4, 3, 5, 10, 8, 9, 7}

	fmt.Printf("unsorted array %v\n", a)

	bubbleSort(a)
	fmt.Printf("sorted array %v\n", a)

}

func bubbleSort(a []int) {

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}

}
