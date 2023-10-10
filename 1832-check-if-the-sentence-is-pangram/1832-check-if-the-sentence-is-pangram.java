import java.util.Map;
import java.util.HashMap;

class Solution {
    public boolean checkIfPangram(String sentence) {
        HashMap<Character, Boolean> hashMap = new HashMap<>();
        for (char ch : sentence.toCharArray()) {
            hashMap.put(ch, true);
        }
        return hashMap.size() == 26;
    }
}