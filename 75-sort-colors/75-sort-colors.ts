/**
 Do not return anything, modify nums in-place instead.
 */
function sortColors(nums: number[]): void {

    var bucketArray: number[] = [0, 0, 0];
    nums.forEach(element => {
        switch (element) {
            case 0:
                bucketArray[0]++;
                break;
            case 1:
                bucketArray[1]++;
                break;
            case 2:
                bucketArray[2]++;
                break;
        }
    })

    var pointer: number = 0;
    nums.forEach((element, index, array) => {
        while (bucketArray[pointer] < 1) {
            pointer++;
        }
        array[index] = pointer;
        bucketArray[pointer]--;
    })

};