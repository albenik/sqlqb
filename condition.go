package sqlqb

import (
	"fmt"
)

type Condition struct {
	left     Element
	right    Element
	operator func(b SQLBuilder, left, right Element) string
}

func eq(b SQLBuilder, left, right Element) string {
	return fmt.Sprintf("%s = %s", left.SQL(b), right.SQL(b))
}

func EQ(left, right Element) *Condition {
	return &Condition{
		left:     left,
		right:    right,
		operator: eq,
	}
}

func (c *Condition) SQL(b SQLBuilder) string {
	return c.operator(b, c.left, c.right)
}
