package query

import (
	"reflect"
)

var collectionTypeRegistry = make(map[string]reflect.Type)

/**
 * Mocking a type registry
 */
func RegisterCollectionType(name string, typ interface{}){
	collectionTypeRegistry[name] = reflect.TypeOf(typ);
}

func NewCollection(name string) ICollection {
	v := reflect.New(collectionTypeRegistry[name])
	coll := v.Interface().(ICollection)
	coll.setType(v.Type())
	coll.setName(name);
	return coll
}

type ICollection interface {
	Entity(string) IEntity
	GetType() string
	GetName() string
	Flush() ICollection
	Append(string, ...V) ICollection
	ToString() string
	setType(reflect.Type)
	setName(string)
}

type Collection struct {
	allowed []IEntity
	actual  []IEntity
	typ     reflect.Type
	name    string
}

func (c *Collection) GetName() string {
	return c.name;
}

func (c *Collection) GetType() string{
	return c.typ.String()
}

func (c *Collection) Entity(name string) IEntity {
	ent := newEntity(name)
	c.allowed = append(c.allowed, ent)
	ent.parent(c)
	return ent
}

func (c *Collection) Flush() ICollection {
	return c
}

func (c *Collection) Append(entName string, attributes...V) ICollection{
	for _, e := range c.allowed {
		if e.GetName() == entName {
			instance := newEntity(entName);
			for _, allowed := range e.getAllowed(){
				instance.Attribute(allowed.GetName());
			}
			for _,v := range attributes{
				instance.Append(v.Name, v.Value)
			}
			c.actual = append(c.actual, instance)
		}
	}

	return c
}

func (c *Collection) ToString() string {
	rtn := c.name
	rtn += " contains : \n"
	for _, e := range c.actual {
		rtn += e.ToString()
	}
	return rtn
}

func (c *Collection) setType(typ reflect.Type) {
	c.typ = typ
}

func (c *Collection) setName(name string) {
	c.name = name
}