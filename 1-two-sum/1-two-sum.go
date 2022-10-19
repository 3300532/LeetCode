func twoSum(nums []int, target int) []int {
    res := []int{}
    a := 0
    b := 0
    c := 1
    for i := b; i <= len(nums); i++ {
        for j := c; j < len(nums); j++ {
            if nums[a]+nums[j] == target {
                res = append(res, a, j)
                
                return res
            }
        }
        a++
        b = a+1
        c++
    }
    return res
}