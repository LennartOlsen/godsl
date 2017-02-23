package persons

type String struct {
	Name string
	Val string
	Typ string
}

func (a *String) Value() interface{} {
	return a.Val
}

func (a *String) Key() string {
	return a.Name
}

func (a *String) Type() string {
	return a.Typ
}

type Integer64 struct {
	Name string
	Val int64
	Typ string
}

func (a *Integer64) Value() interface{} {
	return a.Val
}

func (a *Integer64) Key() string {
	return a.Name
}

func (a *Integer64) Type() string {
	return a.Typ
}