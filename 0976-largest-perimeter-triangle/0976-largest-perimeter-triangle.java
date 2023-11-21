class Solution {
    public int largestPerimeter(int[] nums) {
        // Sort the array in descending order
        Arrays.sort(nums);
        reverse(nums);

        // Iterate through the array and check for triangles
        for (int i = 0; i < nums.length - 2; i++) {
            if (nums[i] < nums[i + 1] + nums[i + 2]) {
                // If the current three sides can form a triangle, return the perimeter
                return nums[i] + nums[i + 1] + nums[i + 2];
            }
        }

        // No triangle with non-zero area found
        return 0;
    }

    // Helper method to reverse the array in-place
    private static void reverse(int[] arr) {
        int start = 0;
        int end = arr.length - 1;
        while (start < end) {
            int temp = arr[start];
            arr[start] = arr[end];
            arr[end] = temp;
            start++;
            end--;
        }
    }
}