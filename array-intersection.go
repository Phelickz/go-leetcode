/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must appear as many times as it shows in both arrays and you may return the result in any order.
*/

package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {

	// creating an empty map
	b := map[int]int{}

	var ans []int

	for _, v := range nums1 {
		// add the values of nums1 to the map above as keys. use anything as value.
		b[v]++
	}

	fmt.Println(b)

	for _, m := range nums2 {
		//check if an item exist in the map
		if _, ok := b[m]; ok {
			if b[m] > 0 {
				ans = append(ans, m)
				// delete(b, m)
				b[m]--
				fmt.Println(b)
			}
		}
	}

	return ans
}
