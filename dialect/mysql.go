package dialect

import (
	"fmt"
)

type MysqlDialect struct{}

func (d MysqlDialect) QuoteIdentifier(s string) string {
	// not the strconv.Quote(), nor %q,
	// just adding double quotes before and after
	return fmt.Sprintf(`"%s"`, s)
}

func (d *MysqlDialect) BindingPlaceholder(int) string {
	return "?"
}
