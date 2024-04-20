class Solution {
    public int trap(int[] height) {
        int answer = 0;

        int leftIndex = 0;
        int rightIndex = height.length-1;
        int leftMax = height[leftIndex];
        int rightMax = height[rightIndex];

        while (leftIndex < rightIndex) {
            if (leftMax <= rightMax) {
                int waterTrapped = leftMax - height[leftIndex];
                if (waterTrapped > 0) {
                    answer += waterTrapped;
                }
                leftIndex++;
                leftMax = Math.max(leftMax, height[leftIndex]);
            } else {
                int waterTrapped = rightMax - height[rightIndex];
                if (waterTrapped > 0) {
                    answer += waterTrapped;
                }
                rightIndex--;
                rightMax = Math.max(rightMax, height[rightIndex]);
            }
        }

        return answer;
    }
}