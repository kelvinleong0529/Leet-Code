class Solution {
    public int maximumWealth(int[][] accounts) {
        int answer = 0;
        for (int[] account: accounts) {
            int wealth = Solution.wealth(account);
            if (wealth > answer) {
                answer = wealth;
            }
        }
        return answer;
    }

    static public int wealth(int[] account) {
        int wealth = 0;
        for (int money: account) {
            wealth += money;
        }
        return wealth;
    }
}
