class Solution {
    public int countMatches(List<List<String>> items, String ruleKey, String ruleValue) {
        int answer = 0;
        int index = 0;
        switch (ruleKey) {
            case "type":
                index = 0;
                break;
            case "color":
                index = 1;
                break;
            case "name":
                index = 2;
                break;
        }
        for (List<String> item: items) {
            if (item.get(index).equals(ruleValue)) {
                answer++;
            }
        }
        return answer;
    }
}