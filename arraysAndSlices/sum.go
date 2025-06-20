package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(ints ...[]int) []int {
	allSum := []int{}

	for _, slice := range ints {
		sum := Sum(slice)
		allSum = append(allSum, sum)
	}
	return allSum
}

func SumAllTails(ints ...[]int) []int {
	allTailSum := []int{}

	for _, slice := range ints {
		if len(slice) == 0 {
			allTailSum = append(allTailSum, 0)
		} else {
			sum := Sum(slice[1:])
			allTailSum = append(allTailSum, sum)
		}
	}
	return allTailSum
}
