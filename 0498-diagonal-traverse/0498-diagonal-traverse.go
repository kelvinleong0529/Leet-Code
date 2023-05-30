func findDiagonalOrder(mat [][]int) []int {
    ROW := len(mat)
    COL := len(mat[0])
    
    hashMap := make([][]int, (ROW-1)+(COL-1)+1)
    
    for r := range mat {
        for c := range mat[r] {
            if (r+c)%2 == 1 {
                hashMap[r+c] = append(hashMap[r+c], mat[r][c])
            } else {
                hashMap[r+c] = append([]int{mat[r][c]}, hashMap[r+c]...)
            }
        }
    }
    
    result := make([]int, 0)
    for i := range hashMap {
        result = append(result, hashMap[i]...)
    }
    return result
}