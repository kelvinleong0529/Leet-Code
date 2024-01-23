class Solution {
    public int findDuplicate(int[] nums) {
        int[] array = new int[nums.length];
        for (int num: nums) {
            if (array[num-1] == -1) {
                return num;
            }
            array[num-1]--;
        }
        return 0;
    }
}