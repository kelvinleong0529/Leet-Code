class Solution {
public:
    bool isValid(std::string s) {
        std::stack<char> stack;
        std::unordered_map<char, char> mapping = {
            {')', '('},
            {']', '['},
            {'}', '{'}
        };

        for (char char_val: s) {
            if (char_val == '{' || char_val == '(' || char_val == '[') {
                stack.push(char_val);
                continue;
            }

            char expected = mapping[char_val];

            if (stack.empty() || stack.top() != expected) {
                return false;
            }
            stack.pop();
        }

        return stack.empty();
    }
};