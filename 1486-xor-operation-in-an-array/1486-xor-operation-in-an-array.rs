impl Solution {
    pub fn xor_operation(n: i32, start: i32) -> i32 {
        let mut result = 0;

        for i in 0..n {
            let num = start + 2 * i;
            result ^= num;
        }

        result
    }
}