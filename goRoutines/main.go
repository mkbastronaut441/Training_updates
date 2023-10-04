package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	go helloWorld(&wg)
	go byeWorld(&wg)

	wg.Add(2)

	wg.Wait()
}

func helloWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello Wolrd")
}

func byeWorld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Bye Wolrd")
}
