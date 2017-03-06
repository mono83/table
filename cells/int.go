package cells

import (
	"github.com/mono83/table"
	"strconv"
)

// Int table.Cell implementation
type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// Width returns cell contents width
func (i Int) Width() int {
	return len(i.String())
}

// PrintFormat returns integer value, aligned to right
func (i Int) PrintFormat(width int) string {
	return table.TextAlignRight(i.String(), width, i.Width())
}
