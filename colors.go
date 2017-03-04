package table

// Colorize applies ANSI terminal color formatting to provided string
func Colorize(value string, colorCode string) string {
	return "\x1b[" + colorCode + "m" + value + "\x1b[0m"
}
