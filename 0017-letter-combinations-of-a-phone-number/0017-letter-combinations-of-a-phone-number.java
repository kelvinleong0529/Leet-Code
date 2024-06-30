class Solution {

        private static final char[][] KEYPAD = {
                new char[] {},
                new char[] {'a', 'b', 'c'},
                new char[] {'d', 'e', 'f'},
                new char[] {'g', 'h', 'i'},
                new char[] {'j', 'k', 'l'},
                new char[] {'m', 'n', 'o'},
                new char[] {'p', 'q', 'r', 's'},
                new char[] {'t', 'u', 'v'},
                new char[] {'w', 'x', 'y', 'z'},
            };

        public List<String> letterCombinations(String digits) {
            List<String> answer = new ArrayList<>();
            
            if (digits.length() == 0) {
                return answer;
            }
            
            recursive(answer, new StringBuilder(), digits, 0);
            return answer;
        }

        private void recursive(List<String> result, StringBuilder stringBuilder, String digits, int index) {
            if (index == digits.length()) {
                result.add(stringBuilder.toString());
                return;
            }

            char[] letters = KEYPAD[digits.charAt(index) - '1'];

            for (char letter: letters) {
                stringBuilder.append(letter);
                recursive(result, stringBuilder, digits, index+1);
                stringBuilder.deleteCharAt(stringBuilder.length()-1);
            }
        }
    }