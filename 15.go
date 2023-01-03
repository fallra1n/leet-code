package main

import (
	"sort"
)

/* -5 -3 -2 1 2 3 4
*   |->
*	i
*
*	   |->      <-|
*	 left	    right
 */

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {

		// if cur num == old num => skip cur while we don't find new num != old num
		if i != 0 {
			oldVal := nums[i-1]
			for i < len(nums) && nums[i] == oldVal {
				i++
			}
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {

			if nums[i]+nums[left]+nums[right] < 0 {
				left++
				continue
			}

			if nums[i]+nums[left]+nums[right] > 0 {
				right--
				continue
			}

			if nums[i]+nums[left]+nums[right] == 0 {
				lastThree := make([]int, 3)

				// if exist is the last three => save values and checking for a match
				if len(res) > 0 {
					lastThree[0] = res[len(res)-1][0]
					lastThree[1] = res[len(res)-1][1]
					lastThree[2] = res[len(res)-1][2]

					// if cur three != last three
					// because our arr sorted => only cur and last triples of res can match
					if !(nums[i] == lastThree[0] && nums[left] == lastThree[1] && nums[right] == lastThree[2]) {
						res = append(res, []int{nums[i], nums[left], nums[right]})
					}
				} else { // if first three
					res = append(res, []int{nums[i], nums[left], nums[right]})
				}

				left++
				right--
			}

		}

	}

	return res
}
