class Solution {
    public int sumOddLengthSubarrays(int[] arr) {
        int answer = 0;
        int length = 1;
        while (length <= arr.length) {
            for (int i = 0; i < arr.length; i++) {
                if (i+length > arr.length) {
                    break;
                }
                answer += Solution.sumArray(arr, i, i+length);
            } 
            length += 2;
        }
        return answer;
    }

    static public int sumArray(int[] arr, int startIndex, int endIndex) {
        int answer = 0;
        for (int i = startIndex; i < endIndex; i++) {
            answer += arr[i];
        }
        return answer;
    }
}