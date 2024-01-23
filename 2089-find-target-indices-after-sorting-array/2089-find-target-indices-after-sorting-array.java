class Solution {
    public List<Integer> targetIndices(int[] nums, int target) {
        List<Integer> array = new ArrayList<>();
        for (int num : nums) {
            while (true) {
                if (array.size() > num-1) {
                    break;
                }
                array.add(0);
            }
            array.set(num-1, array.get(num-1)+1);
        }

        int count = 0;
        List<Integer> answer = new ArrayList<>();
        for (int i = 0; i < array.size(); i++) {
            if (target -1 == i) {
                for (int j = 0; j < array.get(i); j++) {
                    answer.add(count+j);
                }
                break;
            }
            count += array.get(i);
        }
        return answer;
    }
}