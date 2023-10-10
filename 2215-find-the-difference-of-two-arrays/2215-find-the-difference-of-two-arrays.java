import java.util.List;
import java.util.ArrayList;

class Solution {
    public List<List<Integer>> findDifference(int[] nums1, int[] nums2) {
        HashMap<Integer,Boolean> hashMap1 = arrayToHashMap(nums1);
        HashMap<Integer,Boolean> hashMap2 = arrayToHashMap(nums2);
        HashMap<Integer,Boolean> intersection = hashMapIntersection(hashMap1, hashMap2);

        List<List<Integer>> answer = new ArrayList<>();
        int[][] nums = {nums1, nums2};
        for (int[] num : nums) {
            HashMap<Integer,Boolean> innerHashMap = new HashMap<>();
            for (int n : num) {
                if (!intersection.containsKey(n)) {
                    innerHashMap.put(n, true);
                }
            }
            answer.add(new ArrayList<>(innerHashMap.keySet()));
        }
        return answer;
    }

    static public HashMap<Integer,Boolean> arrayToHashMap(int[] nums) {
        HashMap<Integer,Boolean> hashMap = new HashMap<>();
        for (int num : nums) {
            hashMap.put(num, true);
        }
        return hashMap;
    }

    static public HashMap<Integer,Boolean> hashMapIntersection(HashMap<Integer,Boolean> hashMap1, HashMap<Integer,Boolean> hashMap2) {
        HashMap<Integer,Boolean> hashMap = new HashMap<>();
        for (Map.Entry<Integer, Boolean> entry : hashMap1.entrySet()) {
            Integer key = entry.getKey();
            if (hashMap2.containsKey(key)) {
                hashMap.put(key, true);
            }
        }
        return hashMap;
    }
}