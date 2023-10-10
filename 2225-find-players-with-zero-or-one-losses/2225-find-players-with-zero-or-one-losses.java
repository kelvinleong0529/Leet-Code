import java.util.*;

class Solution {
    public List<List<Integer>> findWinners(int[][] matches) {
        HashMap<Integer,Boolean> winnerHashMap = new HashMap<>();
        HashMap<Integer,Integer> loserHashMap = new HashMap<>();

        for (int[] match : matches) {
            int winner = match[0];
            int loser = match[1];

            Integer loserCount = loserHashMap.getOrDefault(loser, 0);
            loserHashMap.put(loser, loserCount + 1);
            if (winnerHashMap.containsKey(loser)) {
                winnerHashMap.remove(loser);
            }

            if (!loserHashMap.containsKey(winner)) {
                winnerHashMap.put(winner, true);
            }
        }

        List<Integer> winners = new ArrayList<>();
        List<Integer> lostOnce = new ArrayList<>();

        for (HashMap.Entry<Integer,Boolean> entry : winnerHashMap.entrySet()) {
            winners.add(entry.getKey());
        }

        for (HashMap.Entry<Integer,Integer> entry : loserHashMap.entrySet()) {
            if (entry.getValue().equals(1)) {
                lostOnce.add(entry.getKey());
            }
        }

        Collections.sort(winners);
        Collections.sort(lostOnce);

        List<List<Integer>> answer = new ArrayList<>();
        answer.add(winners);
        answer.add(lostOnce);

        return answer;
    }
}