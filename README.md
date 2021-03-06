# Table 

Small, easy to use library, that renders arbitrary data as table in console.


# How it works
 
Library operates with two interfaces - `table.Cell` and `table.Inteface`. First one used to describe cell rendering 
contents while the second describes whole table structure (same as Go native `sort.Interface`)

## Tables

Method `table.Print` works only with `table.Interface`, so you have to create structure, that implements it:

```go
type Interface interface {
	// Returns list of headers
	Headers() []string
	// Iterates over own contents, supplying each data row to provided callback function
	EachRow(func(...Cell))
}
```

Example: 
```go
package main

import (
	"github.com/mono83/table"
	"github.com/mono83/table/cells"
	"io/ioutil"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	table.Print(filesTable(files), table.PrintOptions{
		ColumnSeparator: " ",
		NoHeader:        true,
	})
}

// filesTable is proxy structure, that implements table.Interface
type filesTable []os.FileInfo

func (f filesTable) Headers() []string {
	return []string{
		"Name",
		"Type",
		"Size",
	}
}

func (f filesTable) EachRow(callback func(...table.Cell)) {
	for _, info := range f {
		var nodeType table.Cell
		nodeType = cells.Empty{}
		if info.IsDir() {
			nodeType = cells.ColoredYellowHi(cells.String("dir"))
		}

		callback(
			cells.ColoredString(cells.String(info.Name())),
			nodeType,
			cells.Bytes(info.Size()),
		)
	}
}
```


## Cells 

All cells in table must implement `table.Cell interface`:

```go
type Cell interface {
	// Returns cell width
	Width() int
	// Returns plain string representation of cell contents
	String() string
	// Returns formatted value of cell
	PrintFormat(width int) string
}
```

### Implementations

There are some implementations in `cells` package, most of them can be obtained using simple type casting:

|     Type     | Casts&nbsp;from | Usecase |
|--------------|:----------:|---------|
|`cells.Empty` |*none*      | Empty cells |
|`cells.String`|`string`    | String values |
|`cells.Int`   |`int`       | `int` values |
|`cells.Int64` |`int64`     | `int64` values |
|`cells.Bytes` |`int64`     | Byte size of something. This implementation will format value and append suffixes like Kb, Mb, Gb |

### Formatting adapters

#### Coloring

There are some useful coloring adapters, designed as function, that takes `table.Cell` as argument. Suffix `hi`
stands for high-intensity.

* `cells.ColoredRed` and `cells.ColoredRedHi` - red cell coloring
* `cells.ColoredGreen` and `cells.ColoredGreenHi` - green cell coloring
* `cells.ColoredCyan` and `cells.ColoredCyanHi` - cyan cell coloring
* `cells.ColoredYellow` and `cells.ColoredYellowHi` - yellow cell coloring
* `cells.ColoredBlue` and `cells.ColoredBlueHi` - blue cell coloring
* `cells.ColoredWhite` and `cells.ColoredWhiteHi` - white cell coloring
* `cells.ColoredString` - standard color for strings (currently - green)
* `cells.ColoredInt` - standard color for integers (currently - cyan)

#### Alignment

* `cells.AlignLeft`
* `cells.AlignRight`
* `cells.AlignCenter`

### Cell helpers

* `cells.Sprintf` - shortcut to string cell, built using `fmt.Sprintf`

## Print options

* `Writer` - `io.Writer` to use. If omitted, `os.Stdout` will be chosen
* `ColumnSeparator` - string, used as column separator
* `NoDecoration` - if true, color formatting disabled
* `NoHeader` - if true, table header with column names won't be rendered
* `HeaderBorder` - string, used to render border between header and data
