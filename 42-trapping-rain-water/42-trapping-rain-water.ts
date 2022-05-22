function trap(height: number[]): number {
    var arraylength: number = height.length
    var leftPointer: number = 0
    var leftMax: number = height[leftPointer]
    var rightPointer: number = arraylength - 1
    var rightMax: number = height[rightPointer]
    var ans: number = 0;
    var currentIndex = 0
    while (leftPointer < rightPointer) {

        let min: number = leftMax < rightMax ? leftMax : rightMax;
        let currentValue = min - height[currentIndex]
        ans = currentValue > 0 ? ans + currentValue : ans;

        if (height[leftPointer] < height[rightPointer]) {
            leftPointer++;
            leftMax = leftMax > height[leftPointer] ? leftMax : height[leftPointer];
            currentIndex = leftPointer
        }
        else {
            rightPointer--;
            rightMax = rightMax > height[rightPointer] ? rightMax : height[rightPointer];
            currentIndex = rightPointer
        }
    }
    return ans;
};