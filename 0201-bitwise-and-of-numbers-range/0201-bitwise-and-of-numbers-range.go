func rangeBitwiseAnd(left int, right int) int {
    // Bitwise AND of all numbers in the range
    result := left

    for i := left + 1; i <= right; i++ {
        // Perform bitwise AND operation
        result &= i
        
        // If result becomes zero, no need to continue
        if result == 0 {
            break
        }
    }

    return result
}