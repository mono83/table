package table

import (
	"io"
	"os"
	"strings"
)

// PrintOptions contains options, used when table is printed
type PrintOptions struct {
	Writer          io.Writer
	ColumnSeparator string
	NoDecoration    bool
	NoHeader        bool
	HeaderBorder    string
}

// PrintStandard outputs table into standard output using standard options
func PrintStandard(t Interface) error {
	return Print(t, PrintOptions{ColumnSeparator: " | ", HeaderBorder: "-", Writer: os.Stdout})
}

// Print function outputs table using provided printing options
func Print(t Interface, opt PrintOptions) error {
	// Checking defaults
	if opt.Writer == nil {
		opt.Writer = os.Stdout
	}
	if len(opt.ColumnSeparator) == 0 {
		opt.ColumnSeparator = " | "
	}

	// Compiling table
	compiled, err := Compile(t)
	if err != nil {
		return err
	}

	// Calculating total width
	totalWidth := (len(compiled.Headers) - 1) * len(opt.ColumnSeparator)
	for _, w := range compiled.Widths {
		totalWidth += w
	}

	// Printing headers
	if !opt.NoHeader {
		for i, h := range compiled.Headers {
			if i > 0 {
				opt.Writer.Write([]byte(opt.ColumnSeparator))
			}
			opt.Writer.Write([]byte(TextAlignCenter(h, compiled.Widths[i], len(h))))
		}

		opt.Writer.Write([]byte("\n"))

		if len(opt.HeaderBorder) > 0 {
			opt.Writer.Write([]byte(strings.Repeat(opt.HeaderBorder, totalWidth)))
			opt.Writer.Write([]byte("\n"))
		}
	}

	// Printing
	for _, row := range compiled.Data {
		for i, col := range row {
			if i > 0 {
				opt.Writer.Write([]byte(opt.ColumnSeparator))
			}
			opt.Writer.Write([]byte(actual(!opt.NoDecoration, col).PrintFormat(compiled.Widths[i])))
		}
		opt.Writer.Write([]byte("\n"))
	}

	return nil
}

type decorated interface {
	Cell
	Decorated() Cell
}

// actual recursively inspects cell and decoration mode and returns most suitable cell to display
func actual(decorationEnabled bool, c Cell) Cell {
	if decorationEnabled {
		// Decoration is enabled - using cell as-is
		return c
	}
	cc, ok := c.(decorated)
	if !ok {
		// Provided cell is not decorated - using as-is
		return c
	}

	return actual(false, cc.Decorated())
}
