package cells

import "strings"

// Empty represents empty cell
type Empty struct{}

// Width returns cell contents width
func (Empty) Width() int { return 0 }

// PrintFormat returns empty string with desired width
func (Empty) PrintFormat(w int) string { return strings.Repeat(" ", w) }
