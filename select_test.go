package sqlqb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	qb "github.com/albenik/sqlqb"
	"github.com/albenik/sqlqb/dialect"
)

func TestSelectFrom(t *testing.T) {
	t1 := qb.Table("", "table1")
	t2 := qb.Table("", "table2")
	b := qb.NewFactory(dialect.Postgres())

	sql, args := b.Select(t1.Col("col1").As("c11"), t2.Col("col1").As("c21")).
		From(t1).
		InnerJoin(t2, qb.EQ(t2.Col("col1"), t1.Col("col1"))).
		Where(qb.OR(
			qb.EQ(t1.Col("col1"), qb.Bind("foo")),
			qb.Group(qb.AND(
				qb.EQ(t1.Col("col2"), qb.Bind(1)),
				qb.EQ(t1.Col("col3"), qb.Str("bar")),
			)),
			qb.EQ(t2.Col("col1"), qb.Int(123)),
		)).
		OrderBy(qb.ASC(t1.Col("col1")), qb.DESC(t2.Col("col1"))).
		SQL()

	const raw = `SELECT "t1"."col1" AS "c11", "t2"."col1" AS "c21"` +
		` FROM "table1" AS "t1" INNER JOIN "table2" AS "t2" ON "t2"."col1" = "t1"."col1"` +
		` WHERE "t1"."col1" = $1 OR ("t1"."col2" = $2 AND "t1"."col3" = 'bar') OR "t2"."col1" = 123` +
		` ORDER BY "t1"."col1", "t2"."col1" DESC`
	assert.Equal(t, raw, sql)
	assert.Equal(t, []interface{}{"foo", 1}, args)
}
