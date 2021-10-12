/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) 
{
    const set = new Set()
    for (let n of nums)
    {
        if (!set.has(n))
        {
            if (nums.reduce((a,v)=>(v===n?a+1:a),0) == 1)
                return n;
            set.add(n);
        }
    }
};