package sqlqb

import (
	"strings"
)

func AND(refs ...Element) *LogicalStatement {
	return &LogicalStatement{
		operator: " AND ",
		refs:     refs,
	}
}

func OR(refs ...Element) *LogicalStatement {
	return &LogicalStatement{
		operator: " OR ",
		refs:     refs,
	}
}

type LogicalStatement struct {
	operator string
	refs     []Element
}

func (s *LogicalStatement) SQL(b SQLBuilder) string {
	refs := make([]string, 0, len(s.refs))
	for _, r := range s.refs {
		refs = append(refs, r.SQL(b))
	}
	return strings.Join(refs, s.operator)
}

func Group(elements ...Element) GroupStatement {
	return elements
}

type GroupStatement []Element

func (g GroupStatement) SQL(b SQLBuilder) string {
	sb := new(strings.Builder)
	sb.WriteRune('(')
	for _, idn := range g {
		sb.WriteString(idn.SQL(b))
	}
	sb.WriteRune(')')
	return sb.String()
}
