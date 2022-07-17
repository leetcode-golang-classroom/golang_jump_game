package sol

func canJump(nums []int) bool {
	nLen := len(nums)
	maxPos := nums[0]
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for pos := 0; pos < nLen-1; pos++ {
		maxPos = max(maxPos, pos+nums[pos])
		if maxPos <= pos {
			return false
		}
	}
	return true
}
