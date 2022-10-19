func intToRoman(num int) string {
    nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    roms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    res := ""
    for i := 0; i < len(nums); i++ {
        if num/nums[i] != 0 {
            temp := num/nums[i]
            for j := 0; j < temp; j++ {
                res += roms[i]
                num -= nums[i]
            }
        }
    }
    return res    
}