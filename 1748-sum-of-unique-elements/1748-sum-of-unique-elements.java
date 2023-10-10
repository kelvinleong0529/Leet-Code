import java.util.Map;
import java.util.HashMap;

class Solution {
    public int sumOfUnique(int[] nums) {
        HashMap<Integer,Integer> hashMap = new HashMap<>();
        for (int num : nums) {
            Integer count = hashMap.get(num);
            if (count == null) {
                hashMap.put(num, 1);
            } else {
                hashMap.put(num, count + 1);
            }
        }

        int answer = 0;
        for (Map.Entry<Integer,Integer> entry: hashMap.entrySet()) {
            if (entry.getValue().equals(1)) {
                answer += entry.getKey();
            }
        }
        return answer;
    }
}