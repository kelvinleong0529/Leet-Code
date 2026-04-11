class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        List<List<Integer>> res = new ArrayList<>();
        // 1. Sort the array to use the two-pointer strategy
        Arrays.sort(nums);

        for (int i = 0; i < nums.length - 2; i++) {
            // Skip the same element to avoid duplicate triplets
            if (i > 0 && nums[i] == nums[i - 1]) continue;

            // Optimization: If the smallest number is > 0, no triplet can sum to 0
            if (nums[i] > 0) break;

            int left = i + 1;
            int right = nums.length - 1;

            while (left < right) {
                int sum = nums[i] + nums[left] + nums[right];

                if (sum == 0) {
                    res.add(Arrays.asList(nums[i], nums[left], nums[right]));
                    
                    // Skip duplicate values for the left and right pointers
                    while (left < right && nums[left] == nums[left + 1]) left++;
                    while (left < right && nums[right] == nums[right - 1]) right--;
                    
                    // Move pointers inward after finding a valid triplet
                    left++;
                    right--;
                } else if (sum < 0) {
                    left++; // Sum is too small, move left pointer right
                } else {
                    right--; // Sum is too large, move right pointer left
                }
            }
        }
        return res;
    }
}