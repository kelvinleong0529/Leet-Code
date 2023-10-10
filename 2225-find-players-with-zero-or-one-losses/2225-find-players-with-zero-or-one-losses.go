func findWinners(matches [][]int) [][]int {
	winnerHashMap := make(map[int]bool)
	loserHashMap := make(map[int]int)
	for _, v := range matches {
		winner := v[0]
		loser := v[1]

		// update loser
		// 1. increment loser count in loser hashMap
		// 2. if loser presents in winner hashMap, remove it
		loserHashMap[loser]++
		if _, ok := winnerHashMap[loser]; ok {
			delete(winnerHashMap, loser)
		}

		// update winner
		// 1. if winner does not present in loserHashMap, append it
		if _, ok := loserHashMap[winner]; !ok {
			winnerHashMap[winner] = true
		}
	}

	winners := make([]int, 0)
	lostOnce := make([]int, 0)

	for k := range winnerHashMap {
		winners = append(winners, k)
	}

	for k, v := range loserHashMap {
		if v == 1 {
			lostOnce = append(lostOnce, k)
		}
	}

	sort.Slice(winners, func(i, j int) bool {
		return winners[i] < winners[j]
	})
	sort.Slice(lostOnce, func(i, j int) bool {
		return lostOnce[i] < lostOnce[j]
	})

	return [][]int{winners, lostOnce}
}
