class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> hashMap = new HashMap<>();
        int[] answer = new int[2];
        for (int i = 0; i < nums.length; i++) {
            int difference = target - nums[i];
            if (hashMap.containsKey(difference)) {
                answer[0] = hashMap.get(difference);
                answer[1] = i;
            }
            hashMap.put(nums[i], i);
        }
        return answer;
    }
}