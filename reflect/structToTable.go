package reflect

import (
	"errors"
	"github.com/mono83/table"
	"github.com/mono83/table/cells"
	"reflect"
)

// StructToTable produces table from provided struct by reflection analysis
func StructToTable(x interface{}) (table.Interface, error) {
	if x == nil {
		return nil, errors.New("no struct provided")
	}

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("not a struct")
	}

	rs := &reflectedStruct{total: t.NumField()}
	rs.keys = make([]table.Cell, rs.total)
	rs.vals = make([]table.Cell, rs.total)
	for i := 0; i < rs.total; i++ {
		rs.keys[i] = cells.String(t.Field(i).Name)

		value := v.Field(i).Interface()
		if value == nil {
			rs.vals[0] = cells.Empty{}
		} else {
			rs.vals[i] = cells.Sprintf("%+v", value)
			switch value.(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				rs.vals[i] = cells.ColoredCyanHi(rs.vals[i])
			case float32, float64:
				rs.vals[i] = cells.ColoredBlueHi(rs.vals[i])
			case string:
				rs.vals[i] = cells.ColoredGreenHi(rs.vals[i])
			}
		}
	}

	return rs, nil
}

type reflectedStruct struct {
	total int
	keys  []table.Cell
	vals  []table.Cell
}

func (*reflectedStruct) Headers() []string { return []string{"Field", "Value"} }
func (r *reflectedStruct) EachRow(f func(...table.Cell)) {
	for i := 0; i < r.total; i++ {
		f(r.keys[i], r.vals[i])
	}
}
