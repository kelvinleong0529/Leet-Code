class Solution {
    public List<Boolean> kidsWithCandies(int[] candies, int extraCandies) {
        int maxCandy = Solution.maxCandy(candies);
        List<Boolean> answer = new ArrayList<>();
        for (int candy: candies) {
            answer.add(candy+extraCandies >= maxCandy);
        }
        return answer;
    }

    static public int maxCandy(int[] candies) {
        int maxCandy = 0;
        for (int candy : candies) {
            if (candy >= maxCandy) {
                maxCandy = candy;
            }
        }
        return maxCandy;
    }
}
