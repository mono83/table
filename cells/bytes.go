package cells

import (
	"github.com/mono83/table"
	"math/big"
	"strconv"
)

// Bytes implements table.Cell and used to display byte length data
type Bytes int64

func (b Bytes) String() string {
	f, m := b.chunks()
	if f == 0 {
		return "0 " + m
	} else if f == float64(int64(f)) {
		return strconv.Itoa(int(f)) + " " + m
	} else if f > 99 {
		return big.NewFloat(f).Text('f', 1) + " " + m
	}

	return big.NewFloat(f).Text('f', 2) + " " + m
}

// Width returns cell contents width
func (b Bytes) Width() int {
	return len(b.String())
}

// PrintFormat formats cell output
func (b Bytes) PrintFormat(width int) string {
	return table.TextAlignRight(b.String(), width)
}

func (b Bytes) chunks() (float64, string) {
	f64 := float64(int64(b))

	if f64 < 1000 {
		return f64, "B"
	} else if f64 < 10*1024 {
		return f64 / 1024, "Kb"
	} else if f64 < 1000*1024 {
		return f64 / 1024, "Kb"
	} else if f64 < 10*1024*1024 {
		return f64 / 1024 / 1024, "Mb"
	} else if f64 < 1000*1024*1024 {
		return f64 / 1024 / 1024, "Mb"
	}
	return f64 / 1024 / 1024, "Gb"
}
