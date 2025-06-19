package arrays

func Sum(nums []int) int {
	sum := 0

	// for i := 0; i < 5; i++ {
	// 	sum += nums[i]
	// }
	for _, num := range nums {
		sum += num
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }
	var res []int
	for _, nums := range numbersToSum {
		res = append(res, Sum(nums))
	}
	return res
}
