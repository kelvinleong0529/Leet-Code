class Solution {
    public List<Integer> twoOutOfThree(int[] nums1, int[] nums2, int[] nums3) {
        List<Integer> answer = new ArrayList<>();
        Map<Integer, List<Integer>> map = new HashMap<>();
        int [][] numsList = new int[][] {nums1, nums2, nums3};
        for (int i = 0; i < numsList.length; i++) {
            updateMap(map, i, numsList[i]);
        }
        for (Map.Entry<Integer, List<Integer>> entry: map.entrySet()) {
            List<Integer> frequencyList = entry.getValue();
            if (appearAtLeastTwice(frequencyList)) {
                answer.add(entry.getKey());
            }
        }
        return answer;
    }

    private void updateMap(Map<Integer, List<Integer>> map, int index, int[] nums) {
        for (int num: nums) {
            List<Integer> frequencyList = map.computeIfAbsent(num, k -> new ArrayList<>(Arrays.asList(0,0,0)));
            frequencyList.set(index, frequencyList.get(index)+1);
        }
    }

    private boolean appearAtLeastTwice(List<Integer> list) {
        boolean appeared = false;
        for (int num: list) {
            if (num == 0) {
                if (appeared) {
                    return false;
                }
                appeared = true;
            }
        }
        return true;
    }
}