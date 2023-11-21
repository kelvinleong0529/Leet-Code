class Solution {
    public boolean isAcronym(List<String> words, String s) {
        String answer = "";
        for (String word: words) {
            answer += word.charAt(0);
        }
        return answer.equals(s);
    }
}