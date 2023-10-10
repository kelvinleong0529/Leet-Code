import java.util.Map;
import java.util.HashMap;

class Solution {
    public int countConsistentStrings(String allowed, String[] words) {
        HashMap<Character, Boolean> hashMap = new HashMap<>();
        for (char ch : allowed.toCharArray()) {
            hashMap.put(ch, true);
        }

        int answer = 0;
        for (String word : words) {
            if (Solution.valid(hashMap, word)) {
                answer++;
            }
        }

        return answer;
    }

    static public boolean valid(HashMap<Character,Boolean> hashMap, String word) {
        for (char ch : word.toCharArray()) {
            if (!hashMap.containsKey(ch)) {
                return false;
            }
        }
        return true;
    }
}