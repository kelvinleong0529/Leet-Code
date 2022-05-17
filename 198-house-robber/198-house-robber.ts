function rob(nums: number[]): number {
    var rob1:number = 0;
    var rob2:number = 0;
    for (let element of nums) {
        var result1:number = rob1 + element
        var result2:number = rob2
        var temp:number = result1 > result2? result1 : result2;
        rob1 = rob2;
        rob2 = temp;
    }
    return rob2;
};