package arrays

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Run("Test two sum", func(t *testing.T) {
		arr := []int{2, 7, 11, 15}
		target := 9
		expected := []int{0, 1}
		result := TwoSum(arr, target)
		if result[0] != expected[0] || result[1] != expected[1] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Test two sum - Not found", func(t *testing.T) {
		arr := []int{5, 1, 2, 3, 4}
		target := 77
		expected := []int{-1, -1}
		result := TwoSum(arr, target)
		if result[0] != expected[0] || result[1] != expected[1] {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
