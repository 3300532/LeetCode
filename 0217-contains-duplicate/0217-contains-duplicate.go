func containsDuplicate(nums []int) bool {
    	hashSet := make(map[int]struct{})

	for _, el := range nums {
		_, exists := hashSet[el]
		if exists {
			return true
		}

		hashSet[el] = struct{}{}
	}

	return false
}