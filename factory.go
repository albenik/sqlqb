package sqlqb

import (
	"github.com/albenik/sqlqb/dialect"
)

var (
	_ Dialect = (*dialect.MysqlDialect)(nil)
	_ Dialect = (*dialect.PostgresDialect)(nil)
)

type Dialect interface {
	QuoteIdentifier(s string) string
	BindingPlaceholder(int) string
}

type Factory struct {
	dialect Dialect
}

func NewFactory(d Dialect) *Factory {
	return &Factory{dialect: d}
}
