package arrays

func moveZeroes(nums []int) {
	p := 0
	for q := range len(nums) {
		if nums[p] == 0 && nums[q] != 0 {
			nums[p], nums[q] = nums[q], nums[p]
			p++
		} else if nums[q] == 0 && nums[p] != 0 {
			p = q
		}
	}
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)
}

func maxProfit(prices []int) int {
	min_price := prices[0]
	max_profit := 0

	for p := range prices {
		profit := prices[p] - min_price
		max_profit = max(max_profit, profit)
		min_price = min(min_price, prices[p])
	}
	return max_profit
}

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
