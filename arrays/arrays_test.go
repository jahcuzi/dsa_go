package arrays

import (
	"testing"
)

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
