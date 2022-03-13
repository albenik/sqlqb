package sqlqb

import (
	"fmt"
	"strings"
)

type SelectBuilder struct {
	dialect Dialect
	columns []Element
	from    []*TableName
	joins   []*joinStatement
	where   Element
	having  Element
	orderby []Element
	limit   uint64
	offset  uint64
}

func (f *Factory) Select(cols ...Element) *SelectBuilder {
	b := &SelectBuilder{
		dialect: f.dialect,
		columns: cols,
	}
	return b
}

func (b *SelectBuilder) AppendColumns(cols ...Element) *SelectBuilder {
	b.columns = append(b.columns, cols...)
	return b
}

func (b *SelectBuilder) From(tables ...*TableName) *SelectBuilder {
	b.from = tables
	return b
}

func (b *SelectBuilder) AppendTable(table *TableName) *SelectBuilder {
	b.from = append(b.from, table)
	return b
}

func (b *SelectBuilder) join(join string, table *TableName, on Element) *SelectBuilder {
	b.joins = append(b.joins, &joinStatement{
		join:  join,
		table: table,
		on:    on,
	})
	return b
}

func (b *SelectBuilder) InnerJoin(table *TableName, on Element) *SelectBuilder {
	return b.join("INNER JOIN", table, on)
}

func (b *SelectBuilder) LeftJoin(table *TableName, on Element) *SelectBuilder {
	return b.join("LEFT JOIN", table, on)
}

func (b *SelectBuilder) RightJoin(table *TableName, on Element) *SelectBuilder {
	return b.join("RIGHT JOIN", table, on)
}

func (b *SelectBuilder) Where(where Element) *SelectBuilder {
	b.where = where
	return b
}

func (b *SelectBuilder) Having(having Element) *SelectBuilder {
	b.having = having
	return b
}

func (b *SelectBuilder) OrderBy(stmt ...Element) *SelectBuilder {
	b.orderby = stmt
	return b
}

func (b *SelectBuilder) Limit(i uint64) *SelectBuilder {
	b.limit = i
	return b
}

func (b *SelectBuilder) Offset(i uint64) *SelectBuilder {
	b.offset = i
	return b
}

func (b *SelectBuilder) SQL() (string, []interface{}) {
	return b.sql(&sqlBuilder{
		dialect:  b.dialect,
		tables:   make(map[*TableName]string),
		bindings: nil,
	})
}

func (b *SelectBuilder) sql(sqlb *sqlBuilder) (string, []interface{}) {
	for _, t := range b.from {
		sqlb.registerTable(t)
	}
	for _, j := range b.joins {
		sqlb.registerTable(j.table)
	}

	sb := new(strings.Builder)

	sb.WriteString("SELECT ")
	for i, col := range b.columns {
		if i > 0 {
			sb.WriteString(", ")
		}
		if cr, ok := col.(*TableColumnReference); ok {
			sb.WriteString(sqlb.ColumnReference(cr))
		} else {
			sb.WriteString(col.SQL(sqlb))
		}
	}

	sb.WriteString(" FROM ")
	for i, t := range b.from {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(sqlb.TableReference(t))
	}

	for _, join := range b.joins {
		sb.WriteString(" ")
		sb.WriteString(join.SQL(sqlb))
	}

	if b.where != nil {
		sb.WriteString(" WHERE ")
		sb.WriteString(b.where.SQL(sqlb))
	}

	if b.having != nil {
		sb.WriteString(" HAVING ")
		sb.WriteString(b.having.SQL(sqlb))
	}

	if len(b.orderby) > 0 {
		sb.WriteString(" ORDER BY ")
		for i, ord := range b.orderby {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(ord.SQL(sqlb))
		}
	}

	if b.limit > 0 {
		fmt.Fprintf(sb, " LIMIT %d", b.limit)
	}

	if b.offset > 0 {
		fmt.Fprintf(sb, " OFFSET %d", b.offset)
	}

	return sb.String(), sqlb.bindings
}
