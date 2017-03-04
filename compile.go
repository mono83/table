package table

import (
	"errors"
	"fmt"
)

// CompiledTable contains cells and columns dimensions
type CompiledTable struct {
	Widths  []int
	Headers []string
	Data    [][]Cell
}

// Compile method reads cells from interface and calculates column widths
func Compile(t Interface) (*CompiledTable, error) {
	if t == nil {
		return nil, errors.New("Empty table provided")
	}

	cmp := &CompiledTable{
		Headers: t.Headers(),
		Widths:  make([]int, len(t.Headers())),
	}

	for i, h := range cmp.Headers {
		cmp.Widths[i] = len(h)
	}

	var err error
	r := -1
	t.EachRow(func(cells ...Cell) {
		r++
		if err == nil {
			if len(cells) != len(cmp.Widths) {
				err = fmt.Errorf(
					"Found %d columns at row %d, but expected %d (same as headers)",
					len(cells),
					r,
					len(cmp.Widths),
				)
			} else {
				for i, col := range cells {
					w := col.Width()
					if w > cmp.Widths[i] {
						cmp.Widths[i] = w
					}
				}
			}
		}

		cmp.Data = append(cmp.Data, cells)
	})

	if err != nil {
		return nil, err
	}

	return cmp, nil
}
