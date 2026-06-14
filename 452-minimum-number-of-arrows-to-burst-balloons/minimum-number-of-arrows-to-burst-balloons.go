func findMinArrowShots(points [][]int) int {
	ans := 1

	slices.SortFunc(points, func(a, b []int) int {
		return a[1] - b[1]
	})

	arrow := points[0][1]

	for i := 1; i < len(points); i++ {
		if arrow >= points[i][0] {
			continue
		}

		ans += 1
		arrow = points[i][1]
	}

	return ans
}