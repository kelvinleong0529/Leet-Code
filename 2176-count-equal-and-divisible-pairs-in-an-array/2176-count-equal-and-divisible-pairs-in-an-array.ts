function countPairs(nums: number[], k: number): number {
    const hashMap: Map<number, number[]> = new Map<number, number[]>()
    var ans: number = 0
    
    nums.forEach((value:number, index:number) => {
        if (hashMap.has(value)) {
            hashMap.get(value).push(index)
            hashMap.set(value, hashMap.get(value))
        } else {
            hashMap.set(value, [index])
        }
        if (index % k == 0) {
            ans += hashMap.get(value).length - 1
        } else {
            var indexes: number[] = hashMap.get(value)
            indexes.forEach((v:number) => {
                if (v!=index && v*index % k == 0) {
                    ans++
                }
            })
        }
    })
    return ans
};