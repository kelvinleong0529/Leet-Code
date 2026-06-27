func maxArea(height []int) int {
    ans := 0 

	l, r := 0, len(height) - 1

	for l < r {
		h := min(height[l], height[r])
		length := r - l
		area := h * length
		if area >= ans {
			ans = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}