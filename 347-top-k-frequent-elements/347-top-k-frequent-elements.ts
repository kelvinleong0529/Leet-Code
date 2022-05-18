function topKFrequent(nums: number[], k: number): number[] {
    var numsLength: number = nums.length
    var bucketArray: number[][] = new Array(numsLength).fill([]);
    nums.forEach(element => {
        for (let counter: number = 0; counter < numsLength; counter++) {
            if (bucketArray[counter].includes(element)) {
                bucketArray[counter + 1] = [...bucketArray[counter + 1], element]
                bucketArray[counter] = bucketArray[counter].filter(i => i != element);
                break;
            }
            if (counter == numsLength - 1) {
                bucketArray[0] = [...bucketArray[0], element]
            }
        }
    })

    var ans: number[] = []
    for (let counter = numsLength - 1; k > 0 && counter >= 0; counter--) {
        ans = [...ans, ...bucketArray[counter]]
        k -= bucketArray[counter].length
    }

    return ans;
};