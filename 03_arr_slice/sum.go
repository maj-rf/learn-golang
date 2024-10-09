package arr_slice

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

/*
	s := []int{2, 3, 5, 7, 11, 13}
	Slice the slice to give it zero length.
	s = s[:0]
	Output: len=0 cap=6 []

	Extend its length.
	s = s[:4]
	Output: len=4 cap=6 [2 3 5 7]

	Drop its first two values.
	s = s[2:]
	Output: len=2 cap=4 [5 7]
*/

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}
