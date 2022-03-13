package sqlqb

import (
	"fmt"
)

var (
	_ Element = (*TableColumnReference)(nil)
)

func Table(schema, name string) *TableName {
	return &TableName{
		schema: schema,
		name:   name,
	}
}

type TableName struct {
	schema string
	name   string
}

func (t *TableName) Col(name string) *TableColumnReference {
	return &TableColumnReference{
		table: t,
		name:  name,
	}
}

func (t *TableName) SQL(b SQLBuilder) string {
	if t.schema == "" {
		return b.QuoteIdentifier(t.name)
	}
	return fmt.Sprintf("%s.%s", b.QuoteIdentifier(t.schema), b.QuoteIdentifier(t.name))
}

func (t *TableName) String() string {
	if t.schema == "" {
		return t.name
	}
	return fmt.Sprintf("%s.%s", t.schema, t.name)
}

type TableColumnReference struct {
	table *TableName
	name  string
	alias string
}

func (c *TableColumnReference) As(alias string) *TableColumnReference {
	c.alias = alias
	return c
}

func (c *TableColumnReference) SQL(b SQLBuilder) string {
	return fmt.Sprintf("%s.%s", b.TableAlias(c.table), b.QuoteIdentifier(c.name))
}
