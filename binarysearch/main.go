package main

import "fmt"

func main() {

	a := []int{1, 4, 2, 10, 3}
	key := 10
	index := binarySearch(a, key)

	if index != -1 {
		fmt.Printf("Key is at index %v\n", index)
	} else {
		fmt.Println("Key is not found in the array ", index)
	}

}

func binarySearch(a []int, key int) int {

	//not found

	s := 0
	e := len(a) - 1

	for s <= e {
		mid := s + (e-s)/2
		if a[mid] == key {
			return mid

		} else if a[mid] < key {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}

	return -1

}
