class Solution {
    public int[] productExceptSelf(int[] nums) {
        int[] answer = new int[nums.length];
        answer[0] = 1;

        // forward iteration
        for (int i = 0; i < nums.length -1; i++) {
            answer[i+1] = answer[i] * nums[i];
        }

        // backward iteration
        int backwardMultiplier = nums[nums.length-1];
        for (int i = answer.length-2; i >= 0; i--) {
            answer[i] *= backwardMultiplier;
            backwardMultiplier *= nums[i];
        }

        return answer;
    }
}