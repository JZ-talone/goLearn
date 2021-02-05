package solution

/**
[0,1,1,3,3]
4
*/
func MaxAverag(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}

	maxSum := sum

	for i := 0; i < len(nums)-k; i++ {
		sum = sum + nums[i+k] - nums[i]
		maxSum = max1(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
