import java.util.Map;
import java.util.HashMap;

class Solution {
    public int numJewelsInStones(String jewels, String stones) {
        HashMap<Character,Boolean> hashMap = new HashMap<>();
        int answer = 0;
        for (char jewel : jewels.toCharArray()) {
            hashMap.put(jewel, true);
        }

        for (char stone : stones.toCharArray()) {
            if (hashMap.containsKey(stone)) {
                answer++;
            }
        }
        return answer;
    }
}