class Solution {
    public int getCommon(int[] nums1, int[] nums2) {
        int index1 = 0;
        int index2 = 0;

        while (true) {
            if (index1 >= nums1.length || index2 >= nums2.length) {
                return -1;
            }
            
            int num1 = nums1[index1];
            int num2 = nums2[index2];
            
            if (num1 == num2) {
                return num1;
            } else if (num1 < num2) {
                index1++;
            } else {
                index2++;
            }
        }
    }
}