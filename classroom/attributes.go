package classroom

import (
	dsl "github.com/lennartolsen/godsl/query"
	"strconv"
)

type String struct {
	dsl.Attribute
}

func (a *String) ToString() string {
	return a.GetName() + " " + a.Attribute.GetValue().(string)
}


type Integer64 struct {
	dsl.Attribute
}

func (a *Integer64) ToString() string {
	return a.GetName() + " " + strconv.Itoa(a.GetValue().(int))
}