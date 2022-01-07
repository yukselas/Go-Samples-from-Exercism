package twofer

import "strings"

const pattern = "One for you, one for me."

// ShareWith should have a comment documenting it.

func ShareWith(name string) string {

	// Write some code here to pass the test suite.

	// Then remove all the stock comments.

	// They're here to help you get started but they only clutter a finished solution.

	// If you leave them in, reviewers may protest!

	result := pattern

	if name != "" {

		result = strings.Replace(result, "you", name, -1)

	}

	return result

}
