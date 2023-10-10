import java.util.*;

class Solution {
    public String sortString(String s) {
        HashMap<Integer,Integer> hashMap = new HashMap<>();
        for (char ch : s.toCharArray()) {
            int acsiiValue = charToAscii(ch);
            Integer count = hashMap.getOrDefault(acsiiValue, 0);
            hashMap.put(acsiiValue, count + 1);
        }

        String answer = "";
        while (true) {
            if (hashMap.size() == 0) {
                break;
            }

            List<Integer> ascendingKeys = hashMapKeys(hashMap, true);
            for (Integer key : ascendingKeys) {
                answer += asciiToChar(key);
                Integer count = hashMap.get(key);
                if (count - 1 == 0) {
                    hashMap.remove(key);
                } else {
                    hashMap.put(key, count - 1);
                }
            }

            List<Integer> descendingKeys = hashMapKeys(hashMap, false);
            for (Integer key : descendingKeys) {
                answer += asciiToChar(key);
                Integer count = hashMap.get(key);
                if (count - 1 == 0) {
                    hashMap.remove(key);
                } else {
                    hashMap.put(key, count - 1);
                }
            }

        }
        return answer;
    }

    static public int charToAscii(char ch) {
        return ch - 'a';
    }

    static public char asciiToChar(int ascii) {
        return (char) (ascii + 'a');
    }

    static public List<Integer> hashMapKeys(HashMap<Integer,Integer> hashMap, boolean ascending) {
        List<Integer> answer = new ArrayList<>();
        for (HashMap.Entry<Integer,Integer> entry: hashMap.entrySet()) {
            answer.add(entry.getKey());
        }
        if (ascending) {
            Collections.sort(answer);
        } else {
            Collections.sort(answer, Collections.reverseOrder());
        }
        return answer;
    }
}