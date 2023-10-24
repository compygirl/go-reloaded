package functions

import (
	"strings"
)

func isBeginVowel(letter byte) bool {
	if letter == 'a' || letter == 'A' || letter == 'e' || letter == 'E' || letter == 'i' || letter == 'I' || letter == 'o' || letter == 'O' || letter == 'u' || letter == 'U' || letter == 'h' || letter == 'H' {
		return true
	}
	return false
}

func findFristChar(word string) int {
	for i, l := range word {
		if (l >= 'a' && l <= 'z') || (l >= 'A' && l <= 'Z') {
			return i
		}
		if l >= '0' && l <= '9' {
			return i
		}
	}
	return 0
}

func isLetter(s string, l int) bool {
	for i := 0; i < len(s)-l; i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			return true
		}
	}
	return false
}

func ChangeArticle(arr []string) []string {
	for ind, w := range arr {
		if w == "a" || w == "A" || (strings.HasSuffix(w, "a") && !isLetter(w, 1)) || (strings.HasSuffix(w, "A") && !isLetter(w, 1)) {
			if ind+1 < len(arr) {
				index := findIndexStrAfter(arr, ind+1)
				if index > 0 {
					if isBeginVowel(arr[index][findFristChar(arr[index])]) {
						arr[ind] = w + "n"
					}
				}
			}
		}

		if w == "an" || w == "An" || w == "aN" || w == "AN" || ((strings.HasSuffix(w, "an") || strings.HasSuffix(w, "aN") || strings.HasSuffix(w, "AN") || strings.HasSuffix(w, "An")) && !isLetter(w, 2)) {
			if ind+1 < len(arr) {
				index := findIndexStrAfter(arr, ind+1)
				if !isBeginVowel(arr[index][findFristChar(arr[index])]) {
					arr[ind] = w[0 : len(w)-1]
				}
			}
		}
	}
	return arr
}
