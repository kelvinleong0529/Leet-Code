import java.util.HashMap;
import java.util.Map;

class Solution {
    public int numIdenticalPairs(int[] nums) {
        HashMap<Integer,Integer> hashMap = new HashMap<>();
        int answer = 0;
        for (int num: nums) {
            if (hashMap.containsKey(num)) {
                int count = hashMap.get(num);
                answer += count;
                hashMap.put(num, count+1);
            } else {
                hashMap.put(num,1);
            }
        }
        return answer;
    }
}