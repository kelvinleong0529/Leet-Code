/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	modeMap := make(map[int]map[int]bool)
	frequencyMap := make(map[int]int)
	highestFrequency := 0

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
        
        // list all the variables needed
		value := node.Val
		frequency := frequencyMap[value]
		updatedFrequency := frequency + 1
		if updatedFrequency > highestFrequency {
			highestFrequency = updatedFrequency
		}
		
        // update frequency map
		frequencyMap[value] = updatedFrequency
        
        // update mode map
		if mode, ok := modeMap[frequency]; ok {
			delete(mode, value)
		}
    
		if _, ok := modeMap[updatedFrequency]; !ok {
			modeMap[updatedFrequency] = make(map[int]bool)	
		}
        modeMap[updatedFrequency][value] = true
		
		dfs(node.Left)
		dfs(node.Right)
	}
    
    dfs(root)
	
	mode := modeMap[highestFrequency]
    ans := make([]int, len(mode))
    index := 0
	for k := range mode {
		ans[index] = k
        index++
	}

	return ans
}