package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the ratings for our pizza :")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", input)

}
