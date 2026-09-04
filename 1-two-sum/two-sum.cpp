class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> hashMap;
        
        for (int i = 0; i < nums.size(); ++i) {
            int diff = target - nums[i];
            
            // Check if the complement exists in the map
            if (hashMap.find(diff) != hashMap.end()) {
                return {hashMap[diff], i};
            }
            
            // Store the index of the current number
            hashMap[nums[i]] = i;
        }
        
        return {}; // Return an empty vector if no solution is found
    }
};