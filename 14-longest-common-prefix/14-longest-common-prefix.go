func longestCommonPrefix(strs []string) string {
    if len(strs) == 1 {
        return strs[0]
    }
    pref := ""
    limit := MinStr(strs)
    if limit == 0 {
        return pref
    }
        
    for symb := 0; symb < limit; symb++ {
    
        for word := 0; word < len(strs); word++ {
            if strs[0][symb] != strs[word][symb] {
                 return pref
            }
        }
      
        pref += string(strs[0][symb])
        
    }
    
    return pref
}

func MinStr(strs []string) int {
    min := len(strs[0])
    for _, el := range strs {
        if len(el) < min {
            min = len(el)
        }
    }
    return min
}