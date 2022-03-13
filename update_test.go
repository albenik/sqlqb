package sqlqb_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	qb "github.com/albenik/sqlqb"
	"github.com/albenik/sqlqb/dialect"
)

func TestUpdateValues(t *testing.T) {
	t1 := qb.Table("", "table1")
	b := qb.NewFactory(dialect.Postgres())

	sql, args := b.Update(t1).
		Set(t1.Col("col1"), qb.Bind("foo")).
		Set(t1.Col("col2"), qb.Str("bar")).
		Set(t1.Col("col3"), qb.Bind(123)).
		Where(qb.EQ(t1.Col("col1"), qb.Bind("baz"))).
		Suffix(qb.Returning(t1.Col("col1"))).
		SQL()

	const raw = `UPDATE "table1" AS "t1" SET ("col1", "col2", "col3") = ($1, 'bar', $2)` +
		` WHERE "t1"."col1" = $3` +
		` RETURNING "t1"."col1"`
	assert.Equal(t, raw, sql)
	assert.Equal(t, []interface{}{"foo", 123, "baz"}, args)
}

func TestUpdateSelect(t *testing.T) {
	t1 := qb.Table("", "table1")
	t2 := qb.Table("", "table2")
	b := qb.NewFactory(dialect.Postgres())

	sql, args := b.Update(t1).
		Columns(
			t1.Col("col1"),
			t1.Col("col2"),
			t1.Col("col3"),
		).
		Select(b.
			Select(
				t2.Col("col1"),
				t2.Col("col2"),
				t2.Col("col3"),
			).
			From(t2).
			Where(
				qb.EQ(t2.Col("col1"), qb.Bind("foo")),
			),
		).
		Where(
			qb.EQ(t1.Col("col1"), qb.Bind("bar")),
		).
		SQL()

	const raw = `UPDATE "table1" AS "t1" SET ("col1", "col2", "col3") =` +
		` (SELECT "t2"."col1", "t2"."col2", "t2"."col3" FROM "table2" AS "t2" WHERE "t2"."col1" = $1)` +
		` WHERE "t1"."col1" = $2`
	assert.Equal(t, raw, sql)
	assert.Equal(t, []interface{}{"foo", "bar"}, args)
}
