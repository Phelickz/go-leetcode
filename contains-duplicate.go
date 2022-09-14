/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
*/

package main

func containsDuplicate(nums []int) bool {
	newList := []int{}

	for _, v := range nums {
		for _, c := range newList {
			if c == v {
				return true

			} else {
				continue
			}
		}

		newList = append(newList, v)
	}

	return false
}
