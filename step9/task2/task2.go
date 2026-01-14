package main

import "testing"

// func Contains(numbers []int, target int) bool {
// 	for _, num := range numbers {
// 		if num == target {
// 			return true
// 		}
// 	}
// 	return false
// }

func TestContains(t *testing.T) {
	t.Run("TrueTest", func(t *testing.T) {
		t.Parallel()

		result := Contains([]int{1, 2, 3}, 3)

		expectedResult := true

		if result != expectedResult {
			t.Errorf("Got: %t, but wanted: %t\n", result, expectedResult)
		}
	})

	t.Run("FalseTest", func(t *testing.T) {
		t.Parallel()

		result := Contains([]int{1, 2, 3}, 5)

		expectedResult := false

		if result != expectedResult {
			t.Errorf("Got: %t, but wanted: %t\n", result, expectedResult)
		}
	})
}
