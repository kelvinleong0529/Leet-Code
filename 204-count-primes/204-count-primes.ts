function countPrimes(n: number): number {
    if (n<2) return 0;
    var referenceArray:number[] = new Array(n-1).fill(0);
    referenceArray[0] = 1;
    var limit:number = Math.sqrt(n);
    for (let counter = 2; counter < limit;counter++) {
        var multiplier = 2;
        var multiplierResult:number = counter*multiplier;
        while (multiplierResult < n) {
            referenceArray[multiplierResult-1] = 1;
            multiplier++;
            multiplierResult = counter*multiplier;
        }
    }
        
    return referenceArray.filter(element => element===0).length
};