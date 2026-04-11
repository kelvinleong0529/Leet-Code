func threeSum(nums []int) [][]int {
    slices.Sort(nums)

    ans := make([][]int, 0)

    for i := range nums {

        if (i >= len(nums) - 2) {
            break
        }

        if (i > 0 && nums[i] == nums[i-1]) {
            continue
        }

        if nums[i] > 0 {
            break
        }

        left := i + 1
        right := len(nums) - 1

        for {
            if (left >= right) {
                break
            } 

            sum := nums[i] + nums[left] + nums[right]

            if sum == 0 {
                ans = append(ans, []int {nums[i], nums[left], nums[right]})

                for {
                    if left >= right {
                        break
                    }

                    if nums[left] == nums[left+1] {
                        left++
                    } else {
                        break
                    }
                }

                for {
                    if left >= right {
                        break
                    }

                    if nums[right] == nums[right-1] {
                        right--
                    } else {
                        break
                    }
                }

                left++
                right--
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }

    return ans
}