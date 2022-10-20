func bestHand(ranks []int, suits []byte) string {
    suitHashMap := make(map[byte]int)
    rankHashMap := make(map[int]int)
    
    for i := range suits {
        _, okSuits := suitHashMap[suits[i]]
        _, okRanks := rankHashMap[ranks[i]]
        
        if !okSuits {
            suitHashMap[suits[i]] = 0
        }
        suitHashMap[suits[i]]++
        
        if !okRanks {
            rankHashMap[ranks[i]] = 0
        }
        rankHashMap[ranks[i]]++
    }
    
    var answer string
    var maxSameCards int
    
    for _, v := range suitHashMap {
        if v == 5 {
            return "Flush"
        }
    }
    
    
    for _, v := range rankHashMap {
        fmt.Println(maxSameCards)
        if v > maxSameCards {
            maxSameCards = v
        }
    }
    
    switch maxSameCards {
        case 3,4,5:
            answer = "Three of a Kind"
        case 2:
            answer = "Pair"
        default:
            answer = "High Card"
    }

    return answer
}