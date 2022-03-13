package sqlqb

import (
	"strings"
)

type InserBuilder struct {
	dialect  Dialect
	table    *TableName
	columns  []*TableColumnReference
	values   []Element
	selectb  *SelectBuilder
	suffixes []Element
}

func (f *Factory) Insert(table *TableName, cols ...*TableColumnReference) *InserBuilder {
	b := &InserBuilder{
		dialect: f.dialect,
		table:   table,
		columns: cols,
	}
	return b
}

func (b *InserBuilder) AppendColumns(cols ...*TableColumnReference) *InserBuilder {
	b.columns = append(b.columns, cols...)
	return b
}

func (b *InserBuilder) Values(vals ...Element) *InserBuilder {
	b.values = vals
	return b
}

func (b *InserBuilder) Select(sel *SelectBuilder) *InserBuilder {
	b.selectb = sel
	return b
}

func (b *InserBuilder) Suffix(s ...Element) *InserBuilder {
	b.suffixes = append(b.suffixes, s...)
	return b
}

func (b *InserBuilder) SQL() (string, []interface{}) {
	sqlb := &sqlBuilder{
		dialect:  b.dialect,
		tables:   make(map[*TableName]string),
		bindings: nil,
	}
	sqlb.registerTable(b.table)

	sb := new(strings.Builder)

	sb.WriteString("INSERT INTO ")
	sb.WriteString(sqlb.TableReference(b.table))
	sb.WriteString(" (")
	for i, col := range b.columns {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(sqlb.QuoteIdentifier(col.name))
	}
	sb.WriteString(") ")

	if b.selectb != nil {
		sel, _ := b.selectb.sql(sqlb)
		sb.WriteString(sel)
	} else {
		sb.WriteString("VALUES (")
		for i, val := range b.values {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(val.SQL(sqlb))
		}
		sb.WriteString(")")
	}

	for _, s := range b.suffixes {
		sb.WriteString(" ")
		sb.WriteString(s.SQL(sqlb))
	}

	return sb.String(), sqlb.bindings
}
