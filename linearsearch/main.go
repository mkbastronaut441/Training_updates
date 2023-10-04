package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 7}
	key := 6

	index := linearSearch(a, key)
	if index != -1 {
		fmt.Printf("Key is present at index %v\n", index)
	} else {
		fmt.Printf("key is not in the array %v\n", index)
	}

}

func linearSearch(a []int, key int) int {

	for i := 0; i <= len(a); i++ {
		if a[i] == key {
			return i
		}
	}
	return -1

}
