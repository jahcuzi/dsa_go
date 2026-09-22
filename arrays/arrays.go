package arrays

import (
	"slices"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0, len(nums))
	slices.Sort(nums)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		p, q := i+1, len(nums)-1
		for p < q {
			sum := nums[i] + nums[p] + nums[q]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[p], nums[q]})
				q--
				for p < q && nums[q] == nums[q+1] {
					q--
				}
			}
			if sum < 0 {
				p++
				continue
			}
			if sum > 0 {
				q--
				continue
			}
		}
	}
	return res
}

func merge(intervals [][]int) [][]int {

	slices.SortFunc(intervals, func(a []int, b []int) int {
		return a[0] - b[0]
	})

	cur := intervals[0]
	res := make([][]int, 0, len(intervals))

	for i := 1; i < len(intervals); i++ {
		temp := intervals[i]
		if temp[0] <= cur[1] {
			cur[1] = max(temp[1], cur[1])
			continue
		}

		res = append(res, cur)
		cur = temp
	}
	res = append(res, cur)
	return res
}

func removeDuplicates(nums []int) int {
	p := 0
	for q := range nums {
		if nums[p] != nums[q] {
			p++
			nums[p], nums[q] = nums[q], nums[p]
		}
	}

	return p + 1
}

func generate(numRows int) [][]int {
	res := [][]int{}
	prev_row := []int{1}
	res = append(res, prev_row)
	if numRows == 1 {
		return res
	}

	prev_row = []int{1, 1}
	res = append(res, prev_row)
	if numRows == 2 {
		return res
	}
	for range numRows - 2 {
		cur_row := []int{1}
		for i := 0; i < len(prev_row)-1; i++ {
			val := prev_row[i] + prev_row[i+1]
			cur_row = append(cur_row, val)
		}
		cur_row = append(cur_row, 1)
		res = append(res, cur_row)
		prev_row = cur_row
	}
	return res
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range nums {
		nums[i] *= nums[i]
	}

	p, q := 0, n-1
	for i := n - 1; i >= 0; i-- {
		if nums[p] > nums[q] {
			res[i] = nums[p]
			p++
		} else {
			res[i] = nums[q]
			q--
		}
	}
	return res
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	p, c := 0, 1
	for range n - 1 {
		p, c = c, c+p
	}
	return c
}

func majorityElement(nums []int) int {
	majE, majFreq := nums[0], 0

	for _, e := range nums {
		if e == majE {
			majFreq++
		} else {
			majFreq--
		}
		if majFreq == 0 {
			majE, majFreq = e, 1
		}
	}
	return majE
}

func pivotIndex(nums []int) int {
	leftSum, rightSum := 0, 0

	for _, val := range nums {
		rightSum += val
	}
	for i, val := range nums {
		rightSum -= val
		if leftSum == rightSum {
			return i
		}
		leftSum += val
	}
	return -1
}

func runningSum(nums []int) []int {
	rs := 0
	res := []int{}
	for _, val := range nums {
		rs += val
		res = append(res, rs)
	}
	return res
}

func maxProfit2(prices []int) int {
	//--------^ remove the 2 before submitting
	total_profit, p := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		e := prices[i]
		if e > p {
			total_profit += e - p
		}
		p = e
	}
	return total_profit
}

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
