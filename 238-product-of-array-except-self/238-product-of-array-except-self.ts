function productExceptSelf(nums: number[]): number[] {
    var ansArray: number[] = [1];
    var numsLength: number = nums.length;
    // forward iteration
    for (let i = 0; i < numsLength - 1; i++) {
        let temp: number = ansArray[i] * nums[i]
        ansArray.push(temp)
    }
    // backward iteration
    var backward: number = nums[numsLength - 1]
    for (let i = numsLength - 2; i > -1; i--) {
        ansArray[i] *= backward
        backward *= nums[i]
    }

    return ansArray
};