package cells

import "github.com/mono83/table"

type colored struct {
	real      table.Cell
	colorCode string
}

func (c *colored) Width() int {
	return c.real.Width()
}

func (c *colored) PrintFormat(width int) string {
	return table.Colorize(c.real.PrintFormat(width), c.colorCode)
}

func (c *colored) Decorated() table.Cell {
	return c.real
}

// ColoredString wraps provided cell and return new one with applied ANSO coloring
func ColoredString(cell table.Cell) table.Cell {
	return ColoredGreen(cell)
}

// ColoredInt wraps provided cell and return new one with applied ANSO coloring
func ColoredInt(cell table.Cell) table.Cell {
	return ColoredCyanHi(cell)
}

// ColoredRed wraps provided cell and return new one with applied ANSO coloring
func ColoredRed(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "31;2"}
}

// ColoredRedHi wraps provided cell and return new one with applied ANSO coloring
func ColoredRedHi(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "31;1"}
}

// ColoredGreen wraps provided cell and return new one with applied ANSO coloring
func ColoredGreen(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "32;2"}
}

// ColoredGreenHi wraps provided cell and return new one with applied ANSO coloring
func ColoredGreenHi(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "32;1"}
}

// ColoredYellow wraps provided cell and return new one with applied ANSO coloring
func ColoredYellow(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "33;2"}
}

// ColoredYellowHi wraps provided cell and return new one with applied ANSO coloring
func ColoredYellowHi(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "33;1"}
}

// ColoredBlue wraps provided cell and return new one with applied ANSO coloring
func ColoredBlue(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "34;2"}
}

// ColoredBlueHi wraps provided cell and return new one with applied ANSO coloring
func ColoredBlueHi(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "34;1"}
}

// ColoredMagenta wraps provided cell and return new one with applied ANSO coloring
func ColoredMagenta(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "35;2"}
}

// ColoredMagentaHi wraps provided cell and return new one with applied ANSO coloring
func ColoredMagentaHi(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "35;1"}
}

// ColoredCyan wraps provided cell and return new one with applied ANSO coloring
func ColoredCyan(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "36;2"}
}

// ColoredCyanHi wraps provided cell and return new one with applied ANSO coloring
func ColoredCyanHi(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "36;1"}
}

// ColoredWhite wraps provided cell and return new one with applied ANSO coloring
func ColoredWhite(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "37;2"}
}

// ColoredWhiteHi wraps provided cell and return new one with applied ANSO coloring
func ColoredWhiteHi(cell table.Cell) table.Cell {
	return &colored{real: cell, colorCode: "37;1"}
}
