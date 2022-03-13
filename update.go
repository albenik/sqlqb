package sqlqb

import (
	"strings"
)

type UpdateBuilder struct {
	dialect  Dialect
	table    *TableName
	columns  []*TableColumnReference
	values   []Element
	selectb  *SelectBuilder
	where    Element
	suffixes []Element
}

func (f *Factory) Update(table *TableName) *UpdateBuilder {
	b := &UpdateBuilder{
		dialect: f.dialect,
		table:   table,
	}
	return b
}

func (b *UpdateBuilder) Set(col *TableColumnReference, val Element) *UpdateBuilder {
	b.columns = append(b.columns, col)
	b.values = append(b.values, val)
	return b
}

func (b *UpdateBuilder) Columns(cols ...*TableColumnReference) *UpdateBuilder {
	b.columns = append(b.columns, cols...)
	return b
}

func (b *UpdateBuilder) Select(sel *SelectBuilder) *UpdateBuilder {
	b.selectb = sel
	return b
}

func (b *UpdateBuilder) Where(where Element) *UpdateBuilder {
	b.where = where
	return b
}

func (b *UpdateBuilder) Suffix(s ...Element) *UpdateBuilder {
	b.suffixes = append(b.suffixes, s...)
	return b
}

func (b *UpdateBuilder) SQL() (string, []interface{}) {
	sqlb := &sqlBuilder{
		dialect:  b.dialect,
		tables:   make(map[*TableName]string),
		bindings: nil,
	}
	sqlb.registerTable(b.table)

	sb := new(strings.Builder)

	sb.WriteString("UPDATE ")
	sb.WriteString(sqlb.TableReference(b.table))
	sb.WriteString(" SET (")
	for i, col := range b.columns {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(sqlb.QuoteIdentifier(col.name))
	}

	sb.WriteString(") = (")
	if b.selectb != nil {
		sel, _ := b.selectb.sql(sqlb)
		sb.WriteString(sel)
	} else {
		for i, val := range b.values {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(val.SQL(sqlb))
		}
	}
	sb.WriteString(")")

	if b.where != nil {
		sb.WriteString(" WHERE ")
		sb.WriteString(b.where.SQL(sqlb))
	}

	for _, s := range b.suffixes {
		sb.WriteString(" ")
		sb.WriteString(s.SQL(sqlb))
	}

	return sb.String(), sqlb.bindings
}
