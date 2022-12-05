function numIdenticalPairs(nums: number[]): number {
    const hashMap = new Map<number, number>()
    var ans: number = 0
    
    nums.forEach((value:number) => {
        if (hashMap.has(value)) {
            var count: number = hashMap.get(value)
            hashMap.set(value,++count)
            if (count>1) {
                ans += count-1
            }
        } else {
            hashMap.set(value, 1)
        }
    })
    return ans
};