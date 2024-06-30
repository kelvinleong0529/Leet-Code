class Solution {
    public int[] intersection(int[] nums1, int[] nums2) {
        Map<Integer, Boolean> hashMap = new HashMap<>();
        for (int num: nums1) {
            hashMap.put(num, false);
        }
        for (int num: nums2) {
            hashMap.computeIfPresent(num, (k, v) -> true);
        }
        List<Integer> answerList = new ArrayList<>();
        hashMap.forEach((k,v) -> {
            if (v) {
                answerList.add(k);
            }
        });
        int[] ans = new int[answerList.size()];
        for (int i = 0; i < answerList.size(); i++) {
            ans[i] = answerList.get(i);
        }
        return ans;
    }
}