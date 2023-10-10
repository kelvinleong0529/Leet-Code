class Solution {
    public boolean arrayStringsAreEqual(String[] word1, String[] word2) {
        String string1 = Solution.concatStrings(word1);
        String string2 = Solution.concatStrings(word2);
        return string1.equals(string2);
    }

    static public String concatStrings(String[] words) {
        String answer = "";
        for (String word: words) {
            answer += word;
        }
        return answer;
    }
}