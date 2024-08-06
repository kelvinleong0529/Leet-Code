class Solution {
    public int maximumNumberOfStringPairs(String[] words) {
        Map<String, Integer> map = new HashMap<>();
        int answer = 0;
        for (String word: words) {
            String reversedWord = reverseString(word);
            if (map.containsKey(reversedWord)) {
                answer++;
            }
            map.put(word, 1);
        }
        return answer;
    }

    private String reverseString (String str) {
        return new StringBuilder(str).reverse().toString();
    }
}