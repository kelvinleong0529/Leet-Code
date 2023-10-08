class Solution {
    public int[] leftRightDifference(int[] nums) {
        int[] left = new int[nums.length];
        int[] right = new int[nums.length];
        int [] answer = new int[nums.length];

        int total = 0;
        for (int i = 0; i < nums.length; i++) {
            left[i] = total;
            total += nums[i];
        }

        total = 0;
        for (int i = nums.length-1; i >= 0; i--) {
            right[i] = total;
            total += nums[i];
        }

        for (int i = 0; i < nums.length; i++) {
            answer[i] = Math.abs(left[i]-right[i]);
        }

        return answer;
    }
}