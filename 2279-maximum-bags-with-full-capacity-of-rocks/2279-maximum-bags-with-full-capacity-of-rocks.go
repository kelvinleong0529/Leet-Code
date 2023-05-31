func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
    for i := 0; i < len(capacity); i++ {
        capacity[i] -= rocks[i]
    }
    sort.Ints(capacity)
    for i := 0; i < len(capacity); i++ {
        additionalRocks -= capacity[i]
        if additionalRocks == 0 {
            return i + 1
        }
        if additionalRocks < 0 {
            return i
        }
    }
    return len(rocks)
}