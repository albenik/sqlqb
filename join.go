package sqlqb

import (
	"fmt"
)

var _ Element = (*joinStatement)(nil)

type joinStatement struct {
	join  string
	table *TableName
	on    Element
}

func (j *joinStatement) SQL(b SQLBuilder) string {
	return fmt.Sprintf("%s %s ON %s", j.join, b.TableReference(j.table), j.on.SQL(b))
}
