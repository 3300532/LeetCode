func myAtoi(s string) int {
    res := 0
    sign := 1
    signFlag := false
    digit := false
   
    for _, el := range s {

        if (el == '+' || el == '-') && (signFlag || digit) {
            break
        }

        if el == '-' {
            signFlag = true
            sign = -1
            continue
        }

        if el == '+' {
            signFlag = true
            continue
        }

        if (el < '0' || el > '9') && el != ' ' {
            break
        }

        if (digit || signFlag) && el == ' ' {
            break
        }

        if sign == 1 && res > 2147483647 {
            return 2147483647
        }

        if sign == -1 && res < -2147483648 {
            return -2147483648
        }
        
        if el >= '0' && el <= '9' {
            digit = true
            res *= 10
            res += int(el-48)
        }
        
    }

    res *= sign
    
    switch {
        case res < 0 && sign == 1: return 2147483647
        case res > 0 && sign == -1: return -2147483648
        case res < -2147483648: return -2147483648
        case res > 2147483647: return 2147483647
       
    }
    
    return res
}