impl Solution {
    pub fn reverse_words(s: String) -> String {
        let mut ans = String::new();
        let mut temp = String::new();

        for i in s.chars() {
            if i == ' ' {
                ans += &temp.chars().rev().collect::<String>();
                ans += " ";
                temp.clear();
            } else {
                temp.push(i);
            }
        }

        ans += &temp.chars().rev().collect::<String>();

        ans
    }
}