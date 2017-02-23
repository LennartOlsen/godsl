package query

import (
	"reflect"
)

/**
* Entity Specific
 */

var entityTypeRegistry = make(map[string]reflect.Type)

/**
 * Mocking a type registry
 */
func RegisterEntityType(name string, typ interface{}){
	entityTypeRegistry[name] = reflect.TypeOf(typ)
}
func newEntity(name string) IEntity {
	v := reflect.New(entityTypeRegistry[name])
	ent := v.Interface().(IEntity)
	ent.setName(name)
	ent.setType(v.Type())
	return ent
}

type IEntity interface {
	Attribute(typ string) IEntity
	GetType() reflect.Type
	GetName() string
	Entity(name string) IEntity
	Flush() ICollection
	Append(string, interface{})
	ToString() string
	setName(string)
	parent(ICollection)
	getParent() ICollection
	setType(reflect.Type)
	getAllowed() []IAttribute
}

type Ent struct {
	allowed []IAttribute
	actual  []IAttribute
	p       ICollection //parent
	name    string
	typ     reflect.Type
}

func (e *Ent) Attribute(name string) IEntity {
	e.allowed = append(e.allowed, newAttribute(name))
	return e
}

func (e *Ent) GetType() reflect.Type {
	return e.typ
}

func (e *Ent) GetName() string {
	return e.name
}

func (e *Ent) Entity(name string) IEntity {
	return e.p.Entity(name)
}

func (e *Ent) Flush() ICollection {
	return e.p.Flush()
}

/**
 * Adds a new attribute if allowed
 */
func (e *Ent) Append(key string, value interface{}) {
	for _, attr := range e.allowed {
		if(attr.GetName() == key){
			instance := newAttribute(key)
			instance.SetValue(value)
			e.actual = append(e.actual, instance)
		}
	}
}

func (e *Ent) ToString() string {
	rtn := e.GetName() + " \n\t"
	for _, attr := range e.actual {
		rtn += attr.ToString() + " \n\t"
	}
	return rtn
}

func (e *Ent) parent(p ICollection) {
	e.p = p
}

func (e *Ent) getParent() ICollection {
	return e.p
}

func (e *Ent) setName(name string){
	e.name = name
}

func (e *Ent) setType(typ reflect.Type) {
	e.typ = typ
}

func (e *Ent) getAllowed() []IAttribute{
	return e.allowed
}
