impl Solution {
    pub fn restore_string(s: String, indices: Vec<i32>) -> String {
        let mut shuffled_chars: Vec<char> = vec![' '; s.len()];

        for (i, &index) in indices.iter().enumerate() {
            shuffled_chars[index as usize] = s.chars().nth(i).unwrap();
        }

        shuffled_chars.iter().collect()
    }
}