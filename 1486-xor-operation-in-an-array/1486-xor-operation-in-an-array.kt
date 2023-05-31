class Solution {
    fun xorOperation(n: Int, start: Int): Int {
        var result = 0

        for (i in 0 until n) {
            val num = start + 2 * i
            result = result xor num
        }

        return result
    }
}