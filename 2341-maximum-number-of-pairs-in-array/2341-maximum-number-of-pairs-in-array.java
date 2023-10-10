class Solution {
    public int[] numberOfPairs(int[] nums) {
        HashMap<Integer,Integer> hashMap = new HashMap<>();
        for (int num : nums) {
            Integer count = hashMap.getOrDefault(num, 0);
            hashMap.put(num, count + 1);
        }

        int pairCount = 0 ;
        int leftover = 0;
        for (HashMap.Entry<Integer,Integer> entry: hashMap.entrySet()) {
            Integer value = entry.getValue();
            pairCount += value / 2;
            leftover += value % 2;
        }

        return new int[]{pairCount, leftover};
    }
}