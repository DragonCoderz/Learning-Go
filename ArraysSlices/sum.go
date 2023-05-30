// More important information on slices here: https://go.dev/blog/slices-intro
package main

func Sum(numbers []int) int {
	sum := 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	//another way to do it using := range
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// variadic function
func SumAll(numbersToSum ...[]int) []int {
	//lenNumbers := len(numbersToSum)
	//New way to make a Slice
	//sums := make([]int, lenNumbers)
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums
}

//Testing things
