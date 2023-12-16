class Solution {
    public int maximumValue(String[] strs) {
        int answer = 0;
        for (String str: strs) {
            int value = getLength(str);
            if (value > answer) {
                answer = value;
            }
        }
        return answer;
    }

    public int getLength(String str) {
        if (isNumeric(str)) {
            return Integer.parseInt(str);
        } else {
            return str.length();
        }
    }

    public boolean isNumeric(String str) {
        for (int i =0; i <str.length(); i++) {
            if (str.charAt(i) < '0' || str.charAt(i) > '9') {
                return false;
            }
        }
        return true;
    }
}