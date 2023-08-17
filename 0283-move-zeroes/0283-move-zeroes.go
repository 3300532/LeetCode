func moveZeroes(nums []int)  {

    slow := 0
    
    for fast := range nums {
        if nums[slow] == 0 && nums[fast] != 0 {
            nums[slow], nums[fast] = nums[fast], nums[slow]
        }

        if nums[slow] != 0 {
            slow++
        }
    }
}