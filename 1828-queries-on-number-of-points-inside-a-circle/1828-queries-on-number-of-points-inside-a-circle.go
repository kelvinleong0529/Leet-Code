func calculateDistance(x1, x2, y1, y2 int) float64 {
	x := x1 - x2
	y := y1 - y2
	return math.Sqrt(float64(x*x) + float64(y*y))
}

func countPoints(points [][]int, queries [][]int) []int {
	answer := make([]int, len(queries))
	for i, v := range queries {
		x1 := v[0]
		y1 := v[1]
		r := v[2]
		var count int
		for _, point := range points {
			x2 := point[0]
			y2 := point[1]
			if calculateDistance(x1, x2, y1, y2) <= float64(r) {
				count++
			}
		}
		answer[i] = count
	}
	return answer
}
