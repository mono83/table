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
