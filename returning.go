package sqlqb

import (
	"strings"
)

type ReturningStatement struct {
	columns []*TableColumnReference
}

func Returning(cols ...*TableColumnReference) *ReturningStatement {
	return &ReturningStatement{
		columns: cols,
	}
}

func (r *ReturningStatement) SQL(b SQLBuilder) string {
	sb := new(strings.Builder)

	sb.WriteString("RETURNING ")
	for i, col := range r.columns {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(col.SQL(b))
	}

	return sb.String()
}
