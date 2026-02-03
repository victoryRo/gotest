// Package arrays ...
package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func SumAll(groups ...[]int) []int {
	res := make([]int, 0)

	for _, val := range groups {
		res = append(res, Sum(val))
	}
	return res
}

func SumAllTails(groups ...[]int) []int {
	res := make([]int, 0)

	for _, val := range groups {
		if len(val) == 0 {
			val = []int{0}
		}
		res = append(res, Sum(val[1:]))
	}
	return res
}
