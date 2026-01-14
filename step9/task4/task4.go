package main

import (
	"testing"
)

// func AreAnagrams(str1, str2 string) bool {
// 	str1 = strings.ToLower(str1)
// 	str2 = strings.ToLower(str2)

// 	if len(str1) != len(str2) {
// 		return false
// 	}

// 	// Convert strings to slices of runes for sorting
// 	r1 := []rune(str1)
// 	r2 := []rune(str2)

// 	sort.Slice(r1, func(i, j int) bool {
// 		return r1[i] < r1[j]
// 	})

// 	sort.Slice(r2, func(i, j int) bool {
// 		return r2[i] < r2[j]
// 	})

// 	return string(r1) == string(r2)
// }

func TestAreAnagrams(t *testing.T) {
	t.Run("TrueTest", func(t *testing.T) {
		t.Parallel()
		word1 := "Nano"
		word2 := "Anon"
		trueResult := AreAnagrams(word1, word2)
		if !trueResult {
			t.Errorf("Wanted true, but got %t\n", trueResult)
		}
	})

	t.Run("FalseTestLong", func(t *testing.T) {
		t.Parallel()
		word1 := "Nano"
		word2 := "Anonimus"
		falseResult := AreAnagrams(word1, word2)
		if falseResult {
			t.Errorf("Wanted false, but got %t\n", falseResult)
		}
	})

	t.Run("FalseTestShort", func(t *testing.T) {
		t.Parallel()
		word1 := "Nano"
		word2 := "Regi"
		falseResult := AreAnagrams(word1, word2)
		if falseResult {
			t.Errorf("Wanted false, but got %t\n", falseResult)
		}
	})
}
