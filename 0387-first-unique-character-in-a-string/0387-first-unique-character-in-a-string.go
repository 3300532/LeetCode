func firstUniqChar(s string) int {
    strMap := make(map[rune]int)

    for _, el := range s {
        strMap[el]++
    }

    for i, el := range s {
        if strMap[el] == 1 {
            return i
        }
    }

    return -1
}

