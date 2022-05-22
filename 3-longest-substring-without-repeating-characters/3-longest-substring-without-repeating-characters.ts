function lengthOfLongestSubstring(s: string): number {

    var existed: Set<string> = new Set();
    var ans: number = existed.size;
    var stringLength: number = s.length;
    var startIndex: number = 0;
    for (let i = 0; i < stringLength; i++) {
        let element: string = s[i];
        while (existed.has(element)) {
            existed.delete(s[startIndex]);
            startIndex++;
        }
        existed.add(element);
        ans = existed.size > ans ? existed.size : ans
    }
    return ans;
};