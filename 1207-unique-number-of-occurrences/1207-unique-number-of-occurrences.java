import java.util.*;

class Solution {
    public boolean uniqueOccurrences(int[] arr) {
        HashMap<Integer,Integer> hashMap = new HashMap<>();
        for (int v : arr) {
            Integer count = hashMap.getOrDefault(v, 0);
            hashMap.put(v, count + 1);
        }

        HashMap<Integer,Integer> occurenceHashMap = new HashMap<>();
        for (HashMap.Entry<Integer,Integer> entry : hashMap.entrySet()) {
            Integer count = occurenceHashMap.getOrDefault(entry.getValue(), 0);
            occurenceHashMap.put(entry.getValue(), count + 1);
        }

        for (HashMap.Entry<Integer,Integer> entry : occurenceHashMap.entrySet()) {
            if (!entry.getValue().equals(1)) {
                return false;
            }
        }

        return true;
    }
}