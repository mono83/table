package table

import "strings"

// TextAlignLeft inserts certain amount of spaces to the right side of string, increasing it's length to
// expected width.
func TextAlignLeft(s string, expected, current int) string {
	if current >= expected {
		return s
	}

	return s + strings.Repeat(" ", expected-current)
}

// TextAlignRight inserts certain amount of spaces to the left side of string, increasing it's length to
// expected width.
func TextAlignRight(s string, expected, current int) string {
	if current >= expected {
		return s
	}

	return strings.Repeat(" ", expected-current) + s
}

// TextAlignCenter inserts certain amount of spaces to both sides of string, increasing it's length to
// expected width.
func TextAlignCenter(s string, expected, current int) string {
	if current >= expected {
		return s
	}

	left := (expected - current) / 2
	right := expected - current - left

	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}
