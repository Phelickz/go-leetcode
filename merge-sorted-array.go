/*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.*/

package main

import (
	"fmt"
	"sort"
)

// with the sort library
func merge(nums1 []int, m int, nums2 []int, n int) {
	newNums1 := []int{}
	newNums2 := []int{}
	for _, v := range nums1 {
		if v != 0 {
			newNums1 = append(newNums1, v)
		}
	}
	if len(nums2) != 0 {
		for _, v := range nums2 {
			if v != 0 {
				newNums2 = append(newNums2, v)
			}
		}
	}
	newNums1 = append(newNums1, newNums2...)
	if len(newNums1) > 1 {
		sort.Ints(newNums1)
	}
	nums1 = newNums1
	fmt.Println(nums1)
}

//without sort package

// func merge(nums1 []int, m int, nums2 []int, n int)  {
//     i := m - 1
//     j := n - 1
//     for k:=m+n-1; k>=0; k-- {
//         if i>=0 && j>=0 && nums1[i] > nums2[j] {
//             nums1[k] = nums1[i]
//             i--
//         } else if j>=0 {
//             nums1[k] = nums2[j]
//             j--
//         }
//     }
// }
