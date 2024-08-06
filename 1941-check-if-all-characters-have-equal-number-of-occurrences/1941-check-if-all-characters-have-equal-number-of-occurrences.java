class Solution {
    public boolean areOccurrencesEqual(String s) {
        Map<Character, Integer> map = new HashMap<>();
        for (char character: s.toCharArray()) {
            map.put(character, map.getOrDefault(character, 0) + 1);
        }
        int frequency = 0;
        for (Map.Entry<Character, Integer> mapEntry: map.entrySet()) {
            if (frequency == 0) {
                frequency = mapEntry.getValue();
                continue;
            }
            if (mapEntry.getValue() != frequency) {
                return false;
            }
        }
        return true;
    }
}