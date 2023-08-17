func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	hashChanS := make(chan map[byte]int)
	hashChanT := make(chan map[byte]int)

	go prepareHash(s, hashChanS)
	go prepareHash(t, hashChanT)

	hashS := <-hashChanS
	hashT := <-hashChanT

	for key := range hashS {
		if hashS[key] != hashT[key] {
			return false
		}
	}

	return true
}

func prepareHash(input string, hashCh chan map[byte]int) {
	hashRes := make(map[byte]int)

	for i := 0; i < len(input); i++ {
		hashRes[input[i]] += 1
	}

	hashCh <- hashRes
}