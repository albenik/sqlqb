package sqlqb

import (
	"fmt"
	"strings"
)

type Element interface {
	SQL(SQLBuilder) string
}

type SQLBuilder interface {
	QuoteIdentifier(string) string
	TableReference(*TableName) string
	TableAlias(*TableName) string
	ColumnReference(*TableColumnReference) string
	BindVar(interface{}) string
}

type sqlBuilder struct {
	dialect  Dialect
	tables   map[*TableName]string
	bindings []interface{}
}

func (b *sqlBuilder) QuoteIdentifier(s string) string {
	return b.dialect.QuoteIdentifier(s)
}

func (b *sqlBuilder) TableReference(t *TableName) string {
	if alias, ok := b.tables[t]; ok {
		return fmt.Sprintf("%s AS %s", t.SQL(b), alias)
	}

	return t.SQL(b)
}

func (b *sqlBuilder) TableAlias(t *TableName) string {
	alias, _ := b.tables[t]
	return alias
}

func (b *sqlBuilder) ColumnReference(c *TableColumnReference) string {
	sb := new(strings.Builder)

	sb.WriteString(c.SQL(b))
	if c.alias != "" {
		sb.WriteString(" AS ")
		sb.WriteString(b.dialect.QuoteIdentifier(c.alias))
	}

	return sb.String()
}

func (b *sqlBuilder) BindVar(v interface{}) string {
	b.bindings = append(b.bindings, v)
	return b.dialect.BindingPlaceholder(len(b.bindings))
}

func (b *sqlBuilder) registerTable(t *TableName) {
	if _, ok := b.tables[t]; !ok {
		b.tables[t] = b.dialect.QuoteIdentifier(fmt.Sprintf("t%d", len(b.tables)+1))
	}
}
