class Solution {

    /**
     * @param Integer[] $capacity
     * @param Integer[] $rocks
     * @param Integer $additionalRocks
     * @return Integer
     */
    function maximumBags($capacity, $rocks, $additionalRocks) {
        for ($i = 0; $i < count($capacity); $i++) {
            $capacity[$i] -= $rocks[$i];
        }
        sort($capacity);
        for ($i = 0; $i < count($capacity); $i++) {
            $additionalRocks -= $capacity[$i];
            if ($additionalRocks === 0) {
                return $i + 1;
            }
            if ($additionalRocks < 0) {
                return $i;
            }
        }
        return count($rocks);
    }
}