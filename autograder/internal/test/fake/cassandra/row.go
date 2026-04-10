package cassandra

import (
	"github.com/gocql/gocql"
)

// NewRow creates a new gocql.RowData instance.
func NewRow(
	cols []string,
	vals []any,
) gocql.RowData {
	return gocql.RowData{
		Columns: cols,
		Values:  vals,
	}
}
