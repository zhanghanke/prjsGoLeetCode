/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

func twoSum(nums []int, target int) []int {
	var res []int = make([]int, 2)
	if len(nums) <= 1 {
		return res
	}


	var tracker map[int]int = make(map[int]int, len(nums))
	for i, v := range nums {
		tv, ok := tracker[v]
		if !ok {
			tracker[target-v] = i
		} else {
			res[0] = tv
			res[1] = i
			break
		}
	}
	return res
}
