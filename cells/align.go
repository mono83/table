package cells

import "github.com/mono83/table"

type aligned struct {
	alignFunc func(string, int, int) string
	real      table.Cell
}

func (a *aligned) Width() int     { return a.real.Width() }
func (a *aligned) String() string { return a.real.String() }
func (a *aligned) PrintFormat(width int) string {
	return a.alignFunc(a.real.String(), width, a.real.Width())
}

// AlignLeft wraps provided cell and returns new one, where cell output
// will be aligned to left
func AlignLeft(c table.Cell) table.Cell {
	return &aligned{real: c, alignFunc: table.TextAlignLeft}
}

// AlignRight wraps provided cell and returns new one, where cell output
// will be aligned to right
func AlignRight(c table.Cell) table.Cell {
	return &aligned{real: c, alignFunc: table.TextAlignRight}
}

// AlignCenter wraps provided cell and returns new one, where cell output
// will be centered
func AlignCenter(c table.Cell) table.Cell {
	return &aligned{real: c, alignFunc: table.TextAlignCenter}
}
