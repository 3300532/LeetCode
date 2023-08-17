package main

func twoSum(nums []int, target int) []int {
	hashSet := make(map[int]int)

	for i, el := range nums {
		difference := target - el

		_, exists := hashSet[difference]

		if exists {
			return []int{i, hashSet[difference]}
		}
		hashSet[el] = i
	}

	return []int{}
}
