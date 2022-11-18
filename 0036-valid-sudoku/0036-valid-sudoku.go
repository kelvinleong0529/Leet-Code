func transpose(board [][]byte) [][]byte {
    var res = make([][]byte, len(board))
    for i := 0; i < len(board); i++ {
        res[i] = make([]byte, len(board[0]))
        for j := 0; j < len(board[i]); j++ {
            res[i][j] = board[j][i]
        }
    }
    return res
}

func checkRow(row []byte) bool {
    hashMap := make(map[byte]bool)
    for i := range row {
        if string(row[i]) == "." {
            continue
        }
        _, ok := hashMap[row[i]]
        if ok {
            return false
        }
        hashMap[row[i]] = true
    }
    return true
}

func checkBox(board [][]byte) bool {
    hashMap := make(map[byte]bool)
    for i := range board {
        for j := range board[i] {
            if string(board[i][j]) == "." {
                continue
            }
            _, ok := hashMap[board[i][j]]
            if ok {
                return false
            }
            hashMap[board[i][j]] = true
        }
    }
    return true
}

func buildBox(board [][]byte, r, c int) [][]byte {
    var res = make([][]byte, 3)
    var offsetRow = 0
    for i := 0; i < 3; i++{
        res[i] = make([]byte, 3)
        offsetCol := 0
        for j := 0; j < 3; j++ {
            res[i][j] = board[r+offsetRow][c+offsetCol]
            offsetCol++
        }
        offsetRow++
    }
    return res
}

func isValidSudoku(board [][]byte) bool {
    for i := range board {
        if !checkRow(board[i]) {
            return false
        }
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if i%3 ==0 && j%3 == 0 {
                targetBoard := buildBox(board,i,j)
                if !checkBox(targetBoard) {
                    return false
                }
            }
        }
    }
    
    board = transpose(board)
    for i := range board {
        if !checkRow(board[i]) {
            return false
        }
    }
    return true
}