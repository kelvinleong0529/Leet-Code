import java.util.*;

class Solution {
    public int countGoodSubstrings(String s) {
        int answer = 0;
        for (int i = 0; i <= s.length()-3; i++) {
            if (isGoodString(s.substring(i,i+3))) {
                answer++;
            }
        }
        return answer;
    }

    static public boolean isGoodString(String s) {
        HashMap<Character,Boolean> hashMap = new HashMap<>();
        for (char ch : s.toCharArray()) {
            if (hashMap.containsKey(ch)) {
                return false;
            } else {
                hashMap.put(ch, true);
            }
        }
        return true;
    }
}