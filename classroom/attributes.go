package classroom

import (
	"github.com/lennartolsen/godsl/dsl"
	"strconv"
)

type String struct {
	dsl.Attribute
}

func (a *String) ToString() string {
	return a.GetName() + " " + a.Attribute.GetValue().(string)
}

func (a *String) GetPrimitiveType() string {
	return "string";
}


type Integer64 struct {
	dsl.Attribute
}

func (a *Integer64) ToString() string {
	return a.GetName() + " " + strconv.Itoa(a.GetValue().(int))
}

func (a *Integer64) GetPrimitiveType() string {
	return "int64";
}