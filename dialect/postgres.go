package dialect

import (
	"fmt"
)

func Postgres() *PostgresDialect {
	return new(PostgresDialect)
}

type PostgresDialect struct{}

func (d PostgresDialect) QuoteIdentifier(s string) string {
	// not the strconv.Quote(), nor %q,
	// just adding double quotes before and after
	return fmt.Sprintf(`"%s"`, s)
}

func (d *PostgresDialect) BindingPlaceholder(i int) string {
	return fmt.Sprintf("$%d", i)
}
