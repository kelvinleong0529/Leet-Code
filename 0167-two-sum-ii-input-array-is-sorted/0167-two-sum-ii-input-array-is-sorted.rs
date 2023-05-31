impl Solution {
    pub fn two_sum(numbers: Vec<i32>, target: i32) -> Vec<i32> {
        let mut left_pointer = 0;
        let mut right_pointer = numbers.len() - 1;

        loop {
            let temp = numbers[left_pointer] + numbers[right_pointer];

            if temp == target {
                return vec![(left_pointer + 1) as i32, (right_pointer + 1) as i32];
            }

            if temp > target {
                right_pointer -= 1;
            } else {
                left_pointer += 1;
            }
        }
    }
}