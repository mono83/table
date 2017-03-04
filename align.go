package table

import "strings"

// TextAlignLeft inserts certain amount of spaces to the right side of string, increasing it's length to
// expected width.
func TextAlignLeft(s string, w int) string {
	l := len(s)
	if l >= w {
		return s
	}

	return s + strings.Repeat(" ", w-l)
}

// TextAlignRight inserts certain amount of spaces to the left side of string, increasing it's length to
// expected width.
func TextAlignRight(s string, w int) string {
	l := len(s)
	if l >= w {
		return s
	}

	return strings.Repeat(" ", w-l) + s
}

// TextAlignCenter inserts certain amount of spaces to both sides of string, increasing it's length to
// expected width.
func TextAlignCenter(s string, w int) string {
	l := len(s)
	if l >= w {
		return s
	}

	left := (w - l) / 2
	right := w - l - left

	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}
