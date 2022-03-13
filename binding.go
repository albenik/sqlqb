package sqlqb

type BindVar struct {
	variable interface{}
}

func (v *BindVar) SQL(b SQLBuilder) string {
	return b.BindVar(v.variable)
}

func Bind(v interface{}) *BindVar {
	return &BindVar{
		variable: v,
	}
}
