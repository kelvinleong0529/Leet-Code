func majorityElement(nums []int) []int {
    var cand1, cand2, count1, count2 int
    cand1, cand2 = 0, 1 // initialize with different numbers to avoid initial conflict
    count1, count2 = 0, 0

    // 1st pass: find potential candidates
    for _, num := range nums {
        switch {
        case num == cand1:
            count1++
        case num == cand2:
            count2++
        case count1 == 0:
            cand1, count1 = num, 1
        case count2 == 0:
            cand2, count2 = num, 1
        default:
            count1--
            count2--
        }
    }

    // 2nd pass: verify the candidates
    count1, count2 = 0, 0
    for _, num := range nums {
        if num == cand1 {
            count1++
        } else if num == cand2 {
            count2++
        }
    }

    result := []int{}
    n := len(nums)
    if count1 > n/3 {
        result = append(result, cand1)
    }
    if count2 > n/3 {
        result = append(result, cand2)
    }

    return result
}