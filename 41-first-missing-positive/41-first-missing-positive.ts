function firstMissingPositive(nums: number[]): number {

    var numsLength: number = nums.length
    var max: number = numsLength + 1

    nums.forEach((element, index, array) => {
        if (element < 0) array[index] = 0
    })

    for (let i: number = 0; i < nums.length; i++) {
        let temp: number = Math.abs(nums[i])
        if (temp === 0) continue;
        if (temp > nums.length) continue;
        if (nums[temp - 1] === 0) {
            nums[temp - 1] = max * -1;
            continue;
        }
        nums[temp - 1] = Math.abs(nums[temp - 1]) * -1
    }

    for (let i: number = 0; i < nums.length; i++) {
        if (nums[i] > -1) return i + 1;
    }

    return max
};