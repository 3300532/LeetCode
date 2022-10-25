func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    xBase := x
    xRev := 0
    
    for xBase != 0 {
        
        xRev += xBase % 10
        xRev *= 10
        xBase /= 10
    }
    
    if xRev/10 == x {
        return true
    }
    
    return false
}