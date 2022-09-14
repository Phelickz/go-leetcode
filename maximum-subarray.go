/* Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

A subarray is a contiguous part of an array.*/

package main

func maxSubArray(nums []int) int {
	var maxSum int
	var currSum int

	for i := range nums {
		//iterating over the list and adding each time while tracking the current sum and the max sum

		currSum += nums[i]
		if i == 0 {
			maxSum = currSum
		}

		if len(nums) == 1 {
			currSum = nums[0]
			maxSum = currSum
			break
		}

		if currSum > maxSum {
			maxSum = currSum
		}

		if currSum < 0 {
			currSum = 0
		}
	}

	for i := range nums {
		//checking if the any of the single items in the list is greater than the final maxsum. if any is greater, then we pick the item, else the maxsum
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}

	return maxSum
}
