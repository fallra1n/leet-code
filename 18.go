package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	sort.Ints(nums)
	var res [][]int

	for j := 0; j < len(nums)-3; j++ {
		if j != 0 {
			oldVal := nums[j-1]
			for j < len(nums) && nums[j] == oldVal {
				j++
			}
		}

		for i := j + 1; i < len(nums)-2; i++ {

			if i != j+1 {
				oldVal := nums[i-1]
				for i < len(nums) && nums[i] == oldVal {
					i++
				}
			}

			left := i + 1
			right := len(nums) - 1

			for left < right {

				if nums[j]+nums[i]+nums[left]+nums[right] < target {
					left++
					continue
				}

				if nums[j]+nums[i]+nums[left]+nums[right] > target {
					right--
					continue
				}

				if nums[j]+nums[i]+nums[left]+nums[right] == target {
					lastFour := make([]int, 4)

					if len(res) > 0 {
						lastFour[0] = res[len(res)-1][0]
						lastFour[1] = res[len(res)-1][1]
						lastFour[2] = res[len(res)-1][2]
						lastFour[3] = res[len(res)-1][3]

						if !(nums[j] == lastFour[0] && nums[i] == lastFour[1] && nums[left] == lastFour[2] && nums[right] == lastFour[3]) {
							res = append(res, []int{nums[j], nums[i], nums[left], nums[right]})
						}
					} else { // if first three
						res = append(res, []int{nums[j], nums[i], nums[left], nums[right]})
					}

					left++
					right--
				}

			}

		}

	}
	return res
}

//func main() {
//	a := []int{2, 2, 2, 2, 2}
//	fourSum(a, 8)
//}
