package main

// import "fmt"

type NumArray struct {
    nums []int
	prefux_sum []int
}


func Constructor(nums []int) NumArray {
	data := NumArray{nums: nums, prefux_sum: nil}

	data.prefux_sum = append(data.prefux_sum, nums[0])
	for i := 1; i < len(nums); i++ {
		data.prefux_sum = append(data.prefux_sum, nums[i])
		data.prefux_sum[i] += data.prefux_sum[i - 1]
	}

	return data
}


func (this *NumArray) SumRange(left int, right int) int {
	return this.prefux_sum[right] - this.prefux_sum[left] + this.nums[left]
}

// func main() {
// 	data := []int{1, 2, 3, 4}
// 	obj := Constructor(data)
// 	fmt.Println(obj.SumRange(1, 1))
// }