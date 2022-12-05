function countKDifference(nums: number[], k: number): number {
    const hashMap: Map<number, number> = new Map<number, number>()
    var ans: number = 0
    
    nums.forEach((value:number) => {
        var add: number = value+k
        var minus: number = value-k
        if (hashMap.has(add)) {
            ans += hashMap.get(add)
        }
        if (hashMap.has(minus)) {
            ans += hashMap.get(minus)
        }
        if (hashMap.has(value)) {
            hashMap.set(value, hashMap.get(value)+1)
        } else {
            hashMap.set(value, 1)
        }
    })
    return ans
};