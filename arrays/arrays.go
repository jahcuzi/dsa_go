package arrays

func twoSum(arr []int, target int) [2]int {
	idx_map := make(map[int]int)
	for i := range len(arr) {
		if _, found := idx_map[arr[i]]; found {
			return [2]int{idx_map[arr[i]], i}
		}
		idx_map[target-arr[i]] = i

	}
	return [2]int{-1, -1}
}

func maxArea(height []int) int {
	p, q := 0, len(height)-1
	max_area := 0
	for p < q {
		area := min(height[p], height[q]) * (q - p)
		max_area = max(max_area, area)
		if height[p] > height[q] {
			q--
		} else {
			p++
		}
	}
	return max_area
}
