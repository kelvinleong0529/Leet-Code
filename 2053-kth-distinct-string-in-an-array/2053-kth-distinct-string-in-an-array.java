class Solution {
    public String kthDistinct(String[] arr, int k) {
        Map<String, Integer> map = new LinkedHashMap<>();
        for (String str: arr) {
            map.put(str, map.getOrDefault(str, 0) + 1);
        }

        for (Map.Entry<String, Integer> entry: map.entrySet()) {
            String str = entry.getKey();
            int frequency = entry.getValue();
            if (frequency == 1) {
                k--;
                if (k == 0) {
                    return str;
                }
            }
        }
        return "";
    }
}