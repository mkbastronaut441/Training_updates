package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex

func main() {

	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5}
	result := make([]int, len(input))

	for i, data := range input {
		//&result[i] is telling the data that at which position it will be stored
		wg.Add(1)
		go processData(&wg, &result[i], data)
	}
	wg.Wait()

	fmt.Println(result)

}

func processData(wg *sync.WaitGroup, resultDest *int, data int) {

	// lock.Lock()
	defer wg.Done()
	processData := process(data)
	*resultDest = processData
	// lock.Unlock()
}

func process(data int) int {
	return data * 2
}
