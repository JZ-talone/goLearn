package solution

/**
[0,1,1,3,3]
4
*/
func FindMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxSum = max(maxSum, sum)
	}
	return float64(maxSum) / float64(k)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
给你两个长度相同的字符串，s 和 t。

将 s中的第i个字符变到t中的第 i 个字符需要|s[i] - t[i]|的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。

用于变更字符串的最大预算是maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。

如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。

如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。



示例 1：

输入：s = "abcd", t = "bcdf", cost = 3
输出：3
解释：s 中的 "abc" 可以变为 "bcd"。开销为 3，所以最大长度为 3。

示例 2：

输入：s = "abcd", t = "cdef", cost = 3
输出：1
解释：s 中的任一字符要想变成 t 中对应的字符，其开销都是 2。因此，最大长度为 1。
示例 3：

输入：s = "abcd", t = "acde", cost = 0
输出：1
解释：你无法作出任何改动，所以最大长度为 1。

"krrgw"
"zjxss"
19

"krpgjbjjznpzdfy"
"nxargkbydxmsgby"
14
*/
func EqualSubstring(s string, t string, maxCost int) int {
	ls := make([]int, len(s)+1)

	for i, i2 := range s {
		x := int(int32(t[i]) - i2)
		if x < 0 {
			x *= -1
		}
		ls[i] = x
	}

	maxLength := 0
	startIndex := 0
	endIndex := 0
	for endIndex < len(ls) {
		if startIndex == endIndex {
			endIndex++
			continue
		}
		totalCount := 0
		for i := startIndex; i < endIndex; i++ {
			totalCount += ls[i]
		}
		if totalCount <= maxCost {
			maxLength = max(maxLength, endIndex-startIndex)
			endIndex++
			continue
		} else {
			startIndex++
			continue
		}
	}

	return maxLength
}
