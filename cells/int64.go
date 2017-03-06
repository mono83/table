package cells

import (
	"github.com/mono83/table"
	"strconv"
)

// Int64 table.Cell implementation
type Int64 int64

func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

// Width returns cell contents width
func (i Int64) Width() int {
	return len(i.String())
}

// PrintFormat returns integer value, aligned to right
func (i Int64) PrintFormat(width int) string {
	return table.TextAlignRight(i.String(), width, i.Width())
}
