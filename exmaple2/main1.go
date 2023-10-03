package main

import "fmt"

// "sync"

//example 1

// var lock sync.Mutex
// var balance int

// func main() {
// 	var wg sync.WaitGroup

// 	numDeposits := 5

// 	for i := 0; i <= numDeposits; i++ {
// 		wg.Add(1)
// 		go deposit(&wg, 100)
// 	}

// 	wg.Wait()

// 	fmt.Print("Final balance :", balance)
// }

// func deposit(wg *sync.WaitGroup, amount int) {

// 	defer wg.Done()
// 	lock.Lock()
// 	balance += amount
// 	lock.Unlock()
// }

//example 2

// var result int
// var limit = 10

// func main() {

// 	ch := make(chan int)

// 	go counting(ch)
// 	<-ch

// 	fmt.Println("counting done:")

// }

// func counting(ch chan int) {
// 	for i := 1; i <= limit; i++ {
// 		fmt.Println(i)
// 	}
// 	close(ch)

// }

//example 3

// func main() {
// 	out := make(chan int)
// 	in := make(chan int)
// 	results := make(chan int) // New channel to collect results sequentially

// 	var wg sync.WaitGroup

// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go product(in, out, &wg)
// 	}

// 	go func() {
// 		defer close(in)
// 		in <- 1
// 		in <- 2
// 		in <- 3
// 		in <- 4
// 	}()

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()

// 	// Collect results sequentially
// 	go func() {
// 		for result := range out {
// 			results <- result
// 		}
// 		close(results)
// 	}()

// 	// Print results sequentially
// 	for result := range results {
// 		fmt.Println(result)
// 	}
// }

// func product(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for num := range in {
// 		result := num * 2
// 		out <- result
// 	}
// }

//example 4

// func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for job := range jobs {
// 		result := job * 2
// 		results <- result
// 	}
// }

// func main() {
// 	numWorkers := 3
// 	numJobs := 5

// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)
// 	var wg sync.WaitGroup

// 	// Start worker goroutines
// 	for i := 0; i < numWorkers; i++ {
// 		wg.Add(1)
// 		go worker(i, jobs, results, &wg)
// 	}

// 	// Send jobs to workers
// 	for i := 0; i < numJobs; i++ {
// 		jobs <- i
// 	}
// 	close(jobs)

// 	// Collect results
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	// Print results
// 	for result := range results {
// 		fmt.Println("Result:", result)
// 	}
// }

//example 5

// func main() {
// 	out := make(chan int, 3)

// 	go func() {
// 		out <- 1
// 		out <- 2
// 		out <- 3

// 		close(out)
// 	}()

// 	for val := range out {
// 		fmt.Println(val)
// 	}

// }

// example 6

// func main() {
// 	ch := make(chan int, 3)
// 	var wg sync.WaitGroup

// 	wg.Add(2)
// 	go producer(ch, &wg)

// 	go consumer(ch, &wg)

// 	wg.Wait()

// 	fmt.Println("Done")
// }

// func producer(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for val := range ch {
// 		fmt.Println("recived:", val)
// 	}
// }

//example 7

// func main() {
// 	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	ch := make(chan int)

// 	go EvenOrOdd(ch, arr)

// 	for val := range ch {
// 		if val%2 != 0 {
// 			fmt.Println(val, "is odd")
// 		} else {
// 			fmt.Println(val, "is even")
// 		}
// 	}

// }

// func EvenOrOdd(ch chan int, arr []int) {

// 	for i := 0; i <= len(arr); i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

//example 8

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)

	go isPrime(ch, arr)

	for val := range ch {
		if (val % 1) && (val%val) != 0 {
			fmt.Println(val, "is not prime")
		} else {
			fmt.Println(val, "is prime")
		}
	}
}

func isPrime(ch chan int, arr []int) {
	for i := 1; i <= len(arr); i++ {
		ch <- i

	}
	close(ch)
}
