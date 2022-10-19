func romanToInt(s string) int {
    nums := []int{}
    res := 0
    for _, el := range s {
        switch el {
            case 'M': nums = append(nums, 1000)
            case 'D': nums = append(nums, 500)
            case 'C': nums = append(nums, 100)
            case 'L': nums = append(nums, 50)
            case 'X': nums = append(nums, 10)
            case 'V': nums = append(nums, 5)
            case 'I': nums = append(nums, 1)
        }
    }
    
    for i := 0; i < len(nums); i++ {
        
        if i == len(nums)-1 {
            res += nums[i]
        } else {
               if nums[i] < nums[i+1] {
                    res += (nums[i+1]-nums[i])
                    i++
                } else {
                    res += nums[i]
                }
        }
        
     
    }
    return res
}