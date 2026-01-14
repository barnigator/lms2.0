package main

import "testing"

// func ReverseString(input string) string {
// 	runes := []rune(input)
// 	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
// 		runes[i], runes[j] = runes[j], runes[i]
// 	}
// 	return string(runes)
// }

func TestReverseString(t *testing.T) {
	t.Run("Russian", func(t *testing.T) {
		t.Parallel()
		word := "матрешка"
		result := ReverseString(word)
		expectedResult := "акшертам"
		if result != expectedResult {
			t.Errorf("Wnated %s, but got %s\n", expectedResult, result)
		}
	})

	t.Run("English", func(t *testing.T) {
		t.Parallel()
		word := "ball"
		result := ReverseString(word)
		expectedResult := "llab"
		if result != expectedResult {
			t.Errorf("Wnated %s, but got %s\n", expectedResult, result)
		}
	})

	t.Run("Numbers", func(t *testing.T) {
		t.Parallel()
		word := "123456"
		result := ReverseString(word)
		expectedResult := "654321"
		if result != expectedResult {
			t.Errorf("Wnated %s, but got %s\n", expectedResult, result)
		}
	})

	t.Run("Mixed", func(t *testing.T) {
		t.Parallel()
		word := "парFlog2 3@_*"
		result := ReverseString(word)
		expectedResult := "*_@3 2golFрап"
		if result != expectedResult {
			t.Errorf("Wnated %s, but got %s\n", expectedResult, result)
		}
	})
}
