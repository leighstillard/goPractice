package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(sumSlices ...[]int) []int {
	lengthOfNumbers := len(sumSlices)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range sumSlices {
		sums[i] = Sum(numbers)
	}

	return sums
}
