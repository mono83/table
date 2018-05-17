package cells

import (
	"fmt"
	"github.com/mono83/table"
	"time"
)

// Duration represents cell, containing duration
type Duration time.Duration

func (d Duration) String() string {
	td := time.Duration(d)
	if td == 0 {
		return ""
	} else if td < time.Millisecond {
		return fmt.Sprintf("%0.9fs", float64(td.Nanoseconds())/1e9)
	} else if td < time.Second*10 {
		return fmt.Sprintf("%0.3fs", td.Seconds())
	} else if td < time.Second*999 {
		return fmt.Sprintf("%.0fs", td.Seconds())
	} else if td < time.Hour*10 {
		return fmt.Sprintf("%.0fm", td.Minutes())
	}

	return fmt.Sprintf("%.0fh", td.Hours())
}

// Width returns cell contents width
func (d Duration) Width() int { return len(d.String()) }

// PrintFormat returns empty string with desired width
func (d Duration) PrintFormat(w int) string {
	str := d.String()
	return table.TextAlignRight(str, w, len(str))
}
