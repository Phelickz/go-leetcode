/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

package main

func twoSum(nums []int, target int) []int {
	var ans []int
	for i := range nums {
		for m := range nums {
			if m == i {
				break
			}

			sum := nums[i] + nums[m]

			if sum == target {
				ans = append(ans, i, m)
				return ans
			}
		}
	}
	return ans
}

func twoSums(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}
