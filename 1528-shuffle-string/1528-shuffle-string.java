class Solution {
    public String restoreString(String s, int[] indices) {
        char[] shuffled = new char[s.length()];

        for (int i = 0; i < s.length(); i++) {
            shuffled[indices[i]] = s.charAt(i);
        }

        return new String(shuffled);
    }
}