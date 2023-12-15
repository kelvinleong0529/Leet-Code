class Solution {
    public int countAsterisks(String s) {
        int count = 0;
        boolean pair = false;
        for (int i = 0; i < s.length(); i++) {
            switch (s.charAt(i)) {
                case '|':
                    pair = !pair;
                    break;
                case '*':
                    if (!pair) {
                        count++;
                    }
            }
        }
        return count;
    }
}