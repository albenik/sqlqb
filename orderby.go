package sqlqb

import (
	"fmt"
)

var _ Element = (*OrderByStatement)(nil)

func ASC(col *TableColumnReference) *OrderByStatement {
	return &OrderByStatement{col: col}
}

func DESC(col *TableColumnReference) *OrderByStatement {
	return &OrderByStatement{col: col, desc: true}
}

type OrderByStatement struct {
	col  *TableColumnReference
	desc bool
}

func (s *OrderByStatement) SQL(b SQLBuilder) string {
	if s.desc {
		return fmt.Sprintf("%s DESC", s.col.SQL(b))
	}
	return s.col.SQL(b)
}
