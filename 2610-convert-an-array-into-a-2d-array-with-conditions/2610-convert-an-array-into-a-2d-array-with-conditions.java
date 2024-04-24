    class Solution {
        public List<List<Integer>> findMatrix(int[] nums) {
            List<List<Integer>> answer = new ArrayList<>();

            HashMap<Integer, Integer> hashMap = new HashMap<>();
            for (int num : nums) {
                hashMap.compute(num, (k, v) -> v == null? 1 : v+1);
            }

            while (!hashMap.isEmpty()) {
                List<Integer> tempList = new ArrayList<>();
                Iterator<Map.Entry<Integer, Integer>> mapIterator = hashMap.entrySet().iterator();
                while (mapIterator.hasNext()) {
                    Map.Entry<Integer, Integer> entry = mapIterator.next();
                    tempList.add(entry.getKey());
                    if (entry.getValue() != null && entry.getValue() > 1) {
                        hashMap.put(entry.getKey(), entry.getValue()-1);
                    } else {
                        mapIterator.remove();
                    }
                }
                answer.add(tempList);
            }

            return answer;
        }
    }