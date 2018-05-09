package cells

import (
	"github.com/mono83/table"
	"unicode/utf8"
)

// Utf8 table.Cell implementation
type Utf8 string

func (s Utf8) String() string {
	return string(s)
}

// Width returns cell contents width
func (s Utf8) Width() int {
	return utf8.RuneCountInString(string(s))
}

// PrintFormat returns string, aligned to left
func (s Utf8) PrintFormat(width int) string {
	return table.TextAlignLeft(string(s), width, s.Width())
}
