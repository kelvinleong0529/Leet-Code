func canPlaceFlowers(flowerbed []int, n int) bool {
    if len(flowerbed) == 1 {
        if flowerbed[0] == 1 && n > 0 {
            return false
        }
        if n > 2 {
            return false
        } else {
            return true
        }
    }
    for i := range flowerbed {
        if n <= 0 {
            break
        }
        if flowerbed[i] == 1 {
            continue
        }
        if i==0 {
            if flowerbed[i+1] == 0 {
                flowerbed[i] = 1
                n--
            }
        } else if i==len(flowerbed)-1 {
            if flowerbed[i-1] == 0 {
                flowerbed[i] = 1
                n--
            }
        } else if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
            flowerbed[i] = 1
            n--
        }
    }
    if n > 0 {
        return false
    } else {
        return true
    }
}