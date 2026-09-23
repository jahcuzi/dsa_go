package arrays

import (
	"reflect"
	"testing"
)

func TestRotateImage(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		expected := [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}
		rotate(matrix)
		if !reflect.DeepEqual(expected, matrix) {
			t.Errorf("Expected %v, Got %v", expected, matrix)
		}
	})

	t.Run("example 1", func(t *testing.T) {
		matrix := [][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}
		expected := [][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		}
		rotate(matrix)
		if !reflect.DeepEqual(expected, matrix) {
			t.Errorf("Expected %v, Got %v", expected, matrix)
		}
	})
}

func TestContainerWithMostWater(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		expected := 49
		result := maxArea2(height)
		if expected != result {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
	t.Run("Example 2", func(t *testing.T) {
		height := []int{1, 1}
		expected := 1
		result := maxArea2(height)
		if expected != result {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func TestSpiralOrder(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		result := spiralOrder(matrix)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
	t.Run("Example 1", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}
		expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
		result := spiralOrder(matrix)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func TestNextPermutation(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := []int{1, 3, 2}
		nextPermutation(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, Got %v", expected, nums)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{3, 2, 1}
		expected := []int{1, 2, 3}
		nextPermutation(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, Got %v", expected, nums)
		}
	})
	t.Run("Example 3", func(t *testing.T) {
		nums := []int{1, 1, 5}
		expected := []int{1, 5, 1}
		nextPermutation(nums)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Expected %v, Got %v", expected, nums)
		}
	})
}
func TestSubarraySumK(t *testing.T) {

	t.Run("example 1", func(t *testing.T) {
		nums := []int{1, 1, 1}
		k := 2
		expected := 2
		result := subarraySum(nums, k)
		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
	t.Run("example 2", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := 3
		expected := 2
		result := subarraySum(nums, k)
		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
	t.Run("example 3", func(t *testing.T) {
		nums := []int{1}
		k := 0
		expected := 0
		result := subarraySum(nums, k)
		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func TestProductExceptSelf(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		expected := []int{24, 12, 8, 6}
		result := productExceptSelf(arr)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})

	t.Run("example 1", func(t *testing.T) {
		arr := []int{-1, 1, 0, -3, 3}
		expected := []int{0, 0, 9, 0, 0}
		result := productExceptSelf(arr)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func Test3Sum(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		arr := []int{-1, 0, 1, 2, -1, -4}
		expected := [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		}
		result := threeSum(arr)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})

	t.Run("example 2", func(t *testing.T) {
		arr := []int{0, 1, 1}
		expected := [][]int{}
		result := threeSum(arr)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})

	t.Run("example 3", func(t *testing.T) {
		arr := []int{0, 0, 0}
		expected := [][]int{
			{0, 0, 0},
		}
		result := threeSum(arr)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func TestMergeIntervals(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		arr := [][]int{
			{1, 3},
			{2, 6},
			{8, 10},
			{15, 18},
		}
		expected := [][]int{
			{1, 6},
			{8, 10},
			{15, 18},
		}
		result := merge(arr)
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
	t.Run("example 2", func(t *testing.T) {
		arr := [][]int{
			{1, 4},
			{4, 5},
		}
		expected := [][]int{
			{1, 5},
		}
		result := merge(arr)
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
	t.Run("example 3", func(t *testing.T) {
		arr := [][]int{
			{4, 7},
			{1, 4},
		}
		expected := [][]int{
			{1, 7},
		}
		result := merge(arr)
		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func TestRemoveDuplicates(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		arr := []int{1, 1, 2}
		k := 2
		r := removeDuplicates(arr)
		expected := []int{1, 2}
		result := arr
		if k != r {
			t.Errorf("Expected #of distinct elements %v, Got %v", k, r)
		}
		for i := range k {
			if expected[i] != result[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
				return
			}
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		k := 5
		r := removeDuplicates(arr)
		expected := []int{0, 1, 2, 3, 4}
		result := arr
		if k != r {
			t.Errorf("Expected Length %v, Got %v", k, r)
		}
		for i := range k {
			if expected[i] != result[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
				return
			}
		}
	})
}

func TestPascalTriangle(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		n := 5
		expected := [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}
		result := generate(n)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})

	t.Run("Example 1", func(t *testing.T) {
		n := 1
		expected := [][]int{
			{1},
		}
		result := generate(n)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i, row := range result {
			for j := range row {
				if result[i][j] != expected[i][j] {
					t.Errorf("Expected %v, Got %v", expected, result)
				}
			}
		}
	})
}

func TestSortedSquares(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{-4, -1, 0, 3, 10}
		expected := []int{0, 1, 9, 16, 100}
		result := sortedSquares(nums)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})
	t.Run("Example 2", func(t *testing.T) {
		nums := []int{-7, -3, 2, 3, 11}
		expected := []int{4, 9, 9, 49, 121}
		result := sortedSquares(nums)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})
}

func TestFib(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		n := 2
		expected := 1
		result := fib(n)
		if expected != result {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
	t.Run("example 2", func(t *testing.T) {
		n := 3
		expected := 2
		result := fib(n)
		if expected != result {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
	t.Run("example 3", func(t *testing.T) {
		n := 4
		expected := 3
		result := fib(n)
		if expected != result {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func TestMajorityElement(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		arr := []int{3, 2, 3}
		expected := 3
		result := majorityElement(arr)
		if expected != result {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})

	t.Run("Example 1", func(t *testing.T) {
		arr := []int{2, 2, 1, 1, 1, 2, 2}
		expected := 2
		result := majorityElement(arr)
		if expected != result {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func TestPivotIndex(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		arr := []int{1, 7, 3, 6, 5, 6}
		expected := 3
		result := pivotIndex(arr)
		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)

		}
	})

	t.Run("Example 2", func(t *testing.T) {
		arr := []int{1, 2, 3}
		expected := -1
		result := pivotIndex(arr)
		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)

		}
	})

	t.Run("Example 3", func(t *testing.T) {
		arr := []int{2, 1, -1}
		expected := 0
		result := pivotIndex(arr)
		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)

		}
	})
}

func TestRunningSum(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		arr := []int{1, 2, 3, 4}
		expected := []int{1, 3, 6, 10}
		result := runningSum(arr)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})
	t.Run("example 2", func(t *testing.T) {
		arr := []int{1, 1, 1, 1, 1}
		expected := []int{1, 2, 3, 4, 5}
		result := runningSum(arr)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})
	t.Run("example 3", func(t *testing.T) {
		arr := []int{3, 1, 2, 10, 1}
		expected := []int{3, 4, 6, 16, 17}
		result := runningSum(arr)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})
}

func TestMaxProfit2(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}
		expected := 7
		result := maxProfit2(prices)
		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		prices := []int{1, 2, 3, 4, 5}
		expected := 4
		result := maxProfit2(prices)
		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		prices := []int{7, 6, 4, 3, 1}
		expected := 0
		result := maxProfit2(prices)
		if result != expected {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func TestMoveZeroes(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{0, 1, 0, 3, 12}
		expected := []int{1, 3, 12, 0, 0}
		moveZeroes(nums)
		result := nums
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		nums := []int{0}
		expected := []int{0}
		moveZeroes(nums)
		result := nums
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})

}

func TestPlusOne(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		digits := []int{1, 2, 3}
		expected := []int{1, 2, 4}
		result := plusOne(digits)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})
	t.Run("Example 2", func(t *testing.T) {
		digits := []int{4, 3, 2, 1}
		expected := []int{4, 3, 2, 2}
		result := plusOne(digits)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})

	t.Run("Example 3", func(t *testing.T) {
		digits := []int{9}
		expected := []int{1, 0}
		result := plusOne(digits)
		if len(result) != len(expected) {
			t.Errorf("Expected Length %v, Got %v", len(expected), len(result))
			return
		}
		for i := range len(expected) {
			if result[i] != expected[i] {
				t.Errorf("Expected %v, Got %v", expected, result)
			}
		}
	})
}

func TestMaxProfit(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}
		expected := 5
		result := maxProfit(prices)
		if expected != result {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		prices := []int{7, 6, 4, 3, 1}
		expected := 0
		result := maxProfit(prices)
		if expected != result {
			t.Errorf("Expected %v, Got %v", expected, result)
		}
	})
}

func TestTwoSum(t *testing.T) {
	t.Run("Test two sum", func(t *testing.T) {
		arr := []int{2, 7, 11, 15}
		target := 9
		expected := []int{0, 1}
		result := twoSum(arr, target)
		if result[0] != expected[0] || result[1] != expected[1] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Test two sum - Not found", func(t *testing.T) {
		arr := []int{5, 1, 2, 3, 4}
		target := 77
		expected := []int{-1, -1}
		result := twoSum(arr, target)
		if result[0] != expected[0] || result[1] != expected[1] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestMaxArea(t *testing.T) {
	t.Run("Example 1 - Leetcode", func(t *testing.T) {
		arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
		expected := 49
		result := maxArea(arr)
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
	t.Run("Example 2 - Leetcode", func(t *testing.T) {
		arr := []int{1, 1}
		expected := 1
		result := maxArea(arr)
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
