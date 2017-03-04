package table

// Cell represents table cell
type Cell interface {
	// Returns cell width
	Width() int
	// Returns formatted value of cell
	PrintFormat(width int) string
}

// Interface describes table itself
type Interface interface {
	// Returns list of headers
	Headers() []string
	// Iterates over own contents, supplying each data row to provided callback function
	EachRow(func(...Cell))
}
