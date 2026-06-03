func trap(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	lMax, rMax := height[l], height[r]

	for l < r {
		var cur int
		if lMax < rMax {
			l++
			if height[l] > lMax {
				lMax = height[l]
			}
			cur = lMax - height[l]

		} else {
			r--
			if height[r] > rMax {
				rMax = height[r]
			}
			cur = rMax - height[r]
		}
		if cur > 0 {
			ans += cur
		}
	}

	return ans
}
