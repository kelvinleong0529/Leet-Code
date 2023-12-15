class Solution {
    public List<Integer> intersection(int[][] nums) {
        Map<Integer, Integer> hashMap = new HashMap<>();
        for (int[] num : nums) {
            for (int element : num) {
                Integer value = hashMap.getOrDefault(element, 0);
                hashMap.put(element, value+1);
            }
        }
        List<Integer> answer = new ArrayList<>();
        for (Map.Entry<Integer, Integer> entry : hashMap.entrySet()) {
            if (entry.getValue().equals(nums.length)) {
                answer.add(entry.getKey());
            }
        }
        Collections.sort(answer);
        return answer;
    }
}