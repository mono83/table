package cells

import (
	"fmt"
	"github.com/mono83/table"
)

// String table.Cell implementation
type String string

func (s String) String() string {
	return string(s)
}

// Width returns cell contents width
func (s String) Width() int {
	return len(string(s))
}

// PrintFormat returns string, aligned to left
func (s String) PrintFormat(width int) string {
	return table.TextAlignLeft(string(s), width, s.Width())
}

// Sprintf builds string cell using fmt.Sprintf
func Sprintf(format string, a ...interface{}) String {
	return String(fmt.Sprintf(format, a...))
}
