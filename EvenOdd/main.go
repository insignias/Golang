package main

import "fmt"

func main() {
	numberRange := generateRange(10)

	for _, i := range numberRange {
		fmt.Printf("%v is %v\n", i, checkEvenOdd(i))
	}
}
