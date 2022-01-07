package isogram

import "unicode"

func IsIsogram(word string) bool {

	seen := make(map[rune]bool)

	for _, b := range word {

		if !unicode.IsLetter(b) {
			continue

		}

		if _, ok := seen[unicode.ToUpper(b)]; ok {

			return false
		}

		seen[unicode.ToUpper(b)] = true

	}

	return true

}
