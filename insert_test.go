package sqlqb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	qb "github.com/albenik/sqlqb"
	"github.com/albenik/sqlqb/dialect"
)

func TestInsertValues(t *testing.T) {
	t1 := qb.Table("", "table1")
	b := qb.NewFactory(dialect.Postgres())

	sql, args := b.Insert(t1, t1.Col("col1"), t1.Col("col2"), t1.Col("col3")).
		Values(qb.Bind("foo"), qb.Bind(123), qb.Bind(true)).
		Suffix(qb.Returning(t1.Col("col1"))).
		SQL()

	const raw = `INSERT INTO "table1" AS "t1" ("col1","col2","col3") VALUES ($1,$2,$3) RETURNING "t1"."col1"`
	assert.Equal(t, raw, sql)
	assert.Equal(t, []interface{}{"foo", 123, true}, args)
}

func TestInsertSelect(t *testing.T) {
	t1 := qb.Table("", "table1")
	t2 := qb.Table("", "table2")
	b := qb.NewFactory(dialect.Postgres())

	sql, args := b.Insert(t1, t1.Col("col1"), t1.Col("col2"), t1.Col("col3")).
		Select(b.
			Select(t2.Col("col1"), t2.Col("col2"), t2.Col("col3")).
			From(t2).
			Where(qb.EQ(t2.Col("col1"), qb.Bind("foo")))).
		SQL()

	const raw = `INSERT INTO "table1" AS "t1" ("col1","col2","col3")` +
		` SELECT "t2"."col1", "t2"."col2", "t2"."col3"` +
		` FROM "table2" AS "t2"` +
		` WHERE "t2"."col1" = $1`
	assert.Equal(t, raw, sql)
	assert.Equal(t, []interface{}{"foo"}, args)
}
