package arrays

import (
	"testing"
)

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
