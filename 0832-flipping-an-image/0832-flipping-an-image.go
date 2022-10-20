func flipAndInvertImage(image [][]int) [][]int {
    answer := make([][]int,0)
    for i := range image {
        temp := make([]int,0)
        for j := len(image[i])-1; j >= 0; j-- {
            if image[i][j] == 0 {
                temp = append(temp,1)
            } else {
                temp = append(temp,0)
            }
        }
        answer = append(answer,temp)
    }
    return answer
}