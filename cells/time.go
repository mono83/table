package cells

import (
	"github.com/mono83/table"
	"time"
)

// TimeFormatRFC3339 prepares cell with RFC3339 formatting
func TimeFormatRFC3339(t time.Time) TimeFormatCell {
	return TimeFormatCell{Time: t, Layout: time.RFC3339}
}

// TimeFormat builds TimeFormatCell with provided time and layout
func TimeFormat(t time.Time, layout string) TimeFormatCell {
	return TimeFormatCell{Time: t, Layout: layout}
}

// TimeFormatCell used to display time data with required format
type TimeFormatCell struct {
	time.Time
	Layout string
}

func (t TimeFormatCell) String() string {
	return t.Time.Format(t.Layout)
}

// Width returns cell contents width
func (t TimeFormatCell) Width() int {
	return len(t.String())
}

// PrintFormat returns integer value, aligned to right
func (t TimeFormatCell) PrintFormat(width int) string {
	s := t.String()
	return table.TextAlignCenter(s, width, len(s))
}
