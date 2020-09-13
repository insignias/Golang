package main

type numberRange []int

func generateRange(n int) numberRange {
	numberRange := numberRange{}

	for i := 0; i <= n; i++ {
		numberRange = append(numberRange, i)
	}

	return numberRange
}

func checkEvenOdd(i int) string {
	res := ""

	if i%2 == 0 {
		res = "Even"
	} else {
		res = "Odd"
	}

	return res
}
