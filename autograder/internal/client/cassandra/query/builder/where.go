package builder

import "strings"

// Where is a WHERE clause builder for Apache Cassandra.
type Where struct {
	// args holds arguments for the WHERE clause.
	args []any
	// sb holds WHERE clause as a string.
	sb strings.Builder
	// n holds the number of WHERE conditions.
	n int
}

// NewWhere returns new Where instance.
func NewWhere() *Where {
	return &Where{
		args: make([]any, 0),
		sb:   strings.Builder{},
	}
}

// Add adds a new WHERE condition with the given field and value.
func (w *Where) Add(field string, value any) {
	if w.n > 0 {
		w.sb.WriteString(" AND ")
	}
	w.sb.WriteString(field)
	w.sb.WriteString(" = ?")
	w.args = append(w.args, value)
	w.n++
}

// String returns WHERE clause as a string.
func (w *Where) String() string {
	if w.sb.Len() == 0 {
		return ""
	}
	return "WHERE " + w.sb.String()
}

// Args returns arguments for the WHERE clause.
func (w *Where) Args() []any {
	return w.args
}

// Count returns the number of WHERE conditions.
func (w *Where) Count() int {
	return w.n
}
