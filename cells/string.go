package cells

import "github.com/mono83/table"

// String table.Cell implementation
type String string

// Width returns cell contents width
func (s String) Width() int {
	return len(string(s))
}

// PrintFormat returns string, aligned to left
func (s String) PrintFormat(width int) string {
	return table.TextAlignLeft(string(s), width)
}
