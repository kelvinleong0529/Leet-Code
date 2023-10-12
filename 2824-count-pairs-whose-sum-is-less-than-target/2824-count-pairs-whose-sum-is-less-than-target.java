import java.util.*;

class Solution {
    public int countPairs(List<Integer> nums, int target) {
        HashMap<Integer,Integer> hashMap = new HashMap<>();
        int answer = 0;

        for (Integer num : nums) {
            for (HashMap.Entry<Integer,Integer> entry: hashMap.entrySet()) {
                if ( num + entry.getKey() < target) {
                    answer += entry.getValue();
                }
            }
            Integer count = hashMap.getOrDefault(num, 0);
            hashMap.put(num, count+1);
        }

        return answer;
    }
}