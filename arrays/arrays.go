package arrays

func TwoSum(arr []int, target int) [2]int {
	idx_map := make(map[int]int)
	for i := range len(arr) {
		if _, found := idx_map[arr[i]]; found {
			return [2]int{i, idx_map[arr[i]]}
		}
		idx_map[target-arr[i]] = i

	}
	return [2]int{-1, -1}
}
