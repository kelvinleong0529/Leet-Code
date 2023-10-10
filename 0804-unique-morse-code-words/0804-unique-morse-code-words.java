class Solution {
    public static final String[] MORSE = {".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."};

    public int uniqueMorseRepresentations(String[] words) {
        HashMap<String, Boolean> hashMap = new HashMap<>();
        for (String word : words) {
            String morseRepresentation = Solution.getMorseRepresentation(word);
            hashMap.put(morseRepresentation, true);
        }
        return hashMap.size();
    }

    static public String getMorseRepresentation(String word) {
        String answer = "";
        for (char ch: word.toCharArray()) {
            int asciiValue = Solution.charToAscii(ch);
            answer += MORSE[asciiValue];
        }
        return answer;
    }

    static public int charToAscii(char ch) {
        return ch - 'a';
    }
}