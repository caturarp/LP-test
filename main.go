// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("isPalindrome(\"wow\"): %v\n", isPalindrome("wow"))
	fmt.Printf("maxValue([]int{0, -14, 2, 2, 14, 1}): %v\n", maxValue([]int{0, -14, 2, 2, 14, 1}))
	err := printTriangle(5)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	positiveCaseValue, err := calculateFactorialValue(4)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
	fmt.Printf("positive case for calculateFactorialValue: %v\n", positiveCaseValue)
	_, err = calculateFactorialValue(-1)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

}
func isPalindrome(word string) bool {
	var reversedWord string
	for i := range word {
		reversedWord += string(word[len(word)-(i+1)])
	}
	return word == reversedWord
}
func maxValue(numbers []int) int {
	var maxValue int
	for _, number := range numbers {
		if maxValue > number {
			continue
		}
		maxValue = number
	}
	return maxValue
}
func printTriangle(footLength int) error {
	if footLength < 1 {
		return errors.New("input must be more than 0")
	}
	for i := 1; i <= footLength; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
	// for i := 0; i < footLength; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Printf("\n")
	// }
	return nil
}
func calculateFactorialValue(number int) (int, error) {
	if number < 0 {
		return 0, errors.New("factorial of negative numbers is not defined")
	} else if number == 0 || number == 1 {
		return 1, nil
	} else {
		result, _ := calculateFactorialValue(number - 1)
		return number * result, nil
	}

}
