// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	// fmt.Printf("isPalindrome(\"wow\"): %v\n", isPalindrome("wow"))
	// fmt.Printf("maxValue([]int{0, -14, 2, 2, 14, 1}): %v\n", maxValue([]int{0, -14, 2, 2, 14, 1}))
	// err := printTriangle(5)
	// if err != nil {
	// 	fmt.Printf("ERROR: %v\n", err)
	// }
	// positiveCaseValue, err := calculateFactorialValue(4)
	// if err != nil {
	// 	fmt.Printf("ERROR: %v\n", err)
	// 	return
	// }
	// fmt.Printf("positive case for calculateFactorialValue: %v\n", positiveCaseValue)
	// _, err = calculateFactorialValue(-1)
	// if err != nil {
	// 	fmt.Printf("ERROR: %v\n", err)
	// 	return
	// }

	// fmt.Println(nextPalindrome(1000))
	// fmt.Println(fibonacciNumbers(5))
	// fmt.Println(fibonacci(6))
	// fmt.Println(fibonacciSequence(6))
	fmt.Println(nextFibonacciNumber(16)) // expected 21

}

// [0, 1]
// assign to new var left side = current value
// right side arr[index+1]
// condition len-2
//5th 1,1,2,3,5
//

// array fib numbers [0	1	1	2	3	5	8	13	21	34	55	89	144	233	377	610	987	1597	2584	4181	6765]
// input value -> check the next fib number listed
// if 35 -> expeted 55
func nextFibonacciNumber(value int) int {
	fibonacciSequence := fibonacciSequence(value)
	fmt.Println(fibonacciSequence, value)
	for _, number := range fibonacciSequence {
		fmt.Println(number)
		if value < number {
			continue
		}
		return number
	}
	return 1
}

func fibonacciSequence(n int) []int {
	var fibonacciSequence []int
	for i := 1; i < n; i++ {
		fibonacciSequence = append(fibonacciSequence, fibonacci(i))
	}
	return fibonacciSequence
}

// F_{n}=F_{n-1}+F_{n-2}}
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// func fibonacciNumbers(n int) []int {
// 	fibonacciSequence := []int{0, 1}
// 	for i := 0; i < n-1; i++ {
// 		fibonacciSequence = append(fibonacciSequence, fibonacciSequence[i]+fibonacciSequence[i+1])
// 	}
// 	parsedZeroFibonacciSequence := fibonacciSequence[1:]
// 	return parsedZeroFibonacciSequence
// }

// next palindrome
// iterasi biasa
// + the value
// change int to string
// call is palindrome
// if true, return current value
func nextPalindrome(value int) int {
	condition := true
	for condition {
		stringValue := ""
		value += 1
		stringValue = fmt.Sprintf("%v", value)
		palindromeStatus := isPalindrome(stringValue)
		if palindromeStatus {
			return value
		}
	}
	return -1
}

// change arg to int
func isPalindrome(word string) bool {
	var reversedWord string
	for i := range word {
		reversedWord += string(word[len(word)-(i+1)])
	}
	return word == reversedWord
}

// func maxValue(numbers []int) int {
// 	var maxValue int
// 	for _, number := range numbers {
// 		if maxValue > number {
// 			continue
// 		}
// 		maxValue = number
// 	}
// 	return maxValue
// }
// func printTriangle(footLength int) error {
// 	if footLength < 1 {
// 		return errors.New("input must be more than 0")
// 	}
// 	for i := 1; i <= footLength; i++ {
// 		fmt.Println(strings.Repeat("*", i))
// 	}
// 	// for i := 0; i < footLength; i++ {
// 	// 	for j := 0; j <= i; j++ {
// 	// 		fmt.Print("*")
// 	// 	}
// 	// 	fmt.Printf("\n")
// 	// }
// 	return nil
// }
// func calculateFactorialValue(number int) (int, error) {
// 	if number < 0 {
// 		return 0, errors.New("factorial of negative numbers is not defined")
// 	} else if number == 0 || number == 1 {
// 		return 1, nil
// 	} else {
// 		result, _ := calculateFactorialValue(number - 1)
// 		return number * result, nil
// 	}

// }
