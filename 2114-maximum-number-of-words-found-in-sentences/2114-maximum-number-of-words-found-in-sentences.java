class Solution {
    public int mostWordsFound(String[] sentences) {
        int answer = 0;
        for (String sentence: sentences ) {
            String[] words = sentence.split(" ");

            if (words.length > answer) {
                answer = words.length;
            }
        }
        return answer;
    }
}