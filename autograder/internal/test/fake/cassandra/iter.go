package cassandra

import (
	"fmt"
	"reflect"

	"github.com/gocql/gocql"
)

// Iter is a fake implementation of gocql.Iter.
type Iter struct {
	// rows is a list of Cassandra rows to scan.
	rows []gocql.RowData
	// pos it the current position of the iterator.
	pos int
}

// NewIter creates a new Iter instance.
func NewIter(rows ...gocql.RowData) *Iter {
	return &Iter{
		rows: rows,
		pos:  0,
	}
}

// Close implements fake closer.
func (i *Iter) Close() error {
	return nil
}

// Scan scans the next row from the iterator
// into the destinantion values.
func (i *Iter) Scan(dest ...any) bool {
	if i.pos >= len(i.rows) {
		return false
	}
	row := i.rows[i.pos]
	i.pos++
	if len(dest) != len(row.Columns) {
		panic(fmt.Sprintf(
			"mismatched number of args in Scan: expected %d, got %d",
			len(row.Columns), len(dest),
		))
	}
	for idx, val := range row.Values {
		destVal := reflect.ValueOf(dest[idx])
		if destVal.Kind() != reflect.Pointer {
			panic(fmt.Sprintf("dest[%d] is not a pointer", idx))
		}
		elem := destVal.Elem()
		if !elem.CanSet() {
			panic(fmt.Sprintf("dest[%d] is not settable", idx))
		}
		srcVal := reflect.ValueOf(val)
		if srcVal.Type().AssignableTo(elem.Type()) {
			elem.Set(srcVal)
		} else {
			converted := srcVal.Convert(elem.Type())
			elem.Set(converted)
		}
	}
	return true
}
