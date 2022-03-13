package sqlqb

import (
	"fmt"
	"strconv"
)

var _ Element = Value("")

type Value string

func (v Value) SQL(SQLBuilder) string {
	return string(v)
}

func Str(s string) Value {
	return Value(fmt.Sprintf("'%s'", s))
}

func Int(i int) Value {
	return Value(strconv.FormatInt(int64(i), 10))
}

func Bool(b bool) Value {
	return Value(strconv.FormatBool(b))
}
