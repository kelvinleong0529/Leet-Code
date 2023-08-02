func dailyTemperatures(temperatures []int) []int {
	hashMap := make(map[int]int)
	answer := make([]int, len(temperatures))
	queue := make([]int, 0)
	days := make([]int, 0)
	for i, v := range temperatures {
		for {
			if len(queue) == 0 {
				break
			}
			queueLength := len(queue) - 1
			daysLength := len(days) - 1
			temperatureIndex := queue[queueLength]
			element := temperatures[temperatureIndex]
			if v <= element {
				break
			}
			hashMap[temperatureIndex] = days[daysLength]
			queue = queue[:queueLength]
			days = days[:daysLength]
		}
		queue = append(queue, i)
		days = append(days, 0)
		for i, v := range days {
			days[i] = v + 1
		}
	}
	for i := range temperatures {
		answer[i] = hashMap[i]
	}
	return answer
}