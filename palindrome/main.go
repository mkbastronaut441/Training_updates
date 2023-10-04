package main

import "fmt"

func main() {

	s1 := "racecar"
	s2 := ""

	checkPalindrome(s1, s2)

}

func checkPalindrome(s1 string, s2 string) {
	for i := len(s1) - 1; i >= 0; i-- {
		s2 += string(s1[i])

		if s1 != s2 {
			fmt.Printf("string s1 is not a palindrome %v\n", s2)
		}
	}
}
