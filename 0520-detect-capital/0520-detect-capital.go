func detectCapitalUse(word string) bool {
    lower := 0
    upper := 0

    for i := 0; i < len(word); i++ {
        if word[i] >= 'a' && word[i] <= 'z'{
            lower++
        } else if word[i] >= 'A' && word[i] <= 'Z' {
            upper++
        }
    }

    if lower == len(word) {
        return true
    }

    if upper == len(word) {
        return true
    }

    if (word[0]>='A' && word[0]<='Z') && (upper==1) {
        return true
    }

    return false
}