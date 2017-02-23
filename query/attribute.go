package query

import "reflect"

/**
* Attribute specific
 */

/**
 * Mocking a type registry
 */
var attributeTypeRegistry = make(map[string]reflect.Type)
func RegisterAttributeType(name string, typ interface{}){
	attributeTypeRegistry[name] = reflect.TypeOf(typ)
}
func newAttribute(name string) IAttribute {
	v := reflect.New(attributeTypeRegistry[name])
	ent := v.Interface().(IAttribute)
	ent.setName(name)
	ent.setType(v.Type())
	return ent
}


type IAttribute interface{
	GetValue() interface{} /* This is kinda floaty */
	SetValue(interface{})
	GetName() string
	GetType() reflect.Type
	ToString() string
	setType( reflect.Type )
	setName( string )
}

type Attribute struct {
	name string
	v interface{}
	typ reflect.Type
}

func (a *Attribute) SetValue(v interface{}) {
	a.v = v
}

func (a *Attribute) GetValue() interface{} {
	return a.v
}

func (a *Attribute) GetName() string {
	return a.name
}

func (a *Attribute) GetType() reflect.Type {
	return a.typ
}

func (a *Attribute) ToString() string {
	return "BAAD"
}

func (a *Attribute) setType(typ reflect.Type){
	a.typ = typ
}

func (a *Attribute) setName(name string){
	a.name = name
}