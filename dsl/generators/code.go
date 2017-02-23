package generators

import (
	"github.com/lennartolsen/godsl/dsl"
	"strings"
	"fmt"
)

func Code( c dsl.ICollection ) {
	cls := []class{}
	for _, e := range c.GetAllowed() {
		attrs := []attribute{}
		for _, attr := range e.GetAllowed() {
			a := attribute{attr.GetName(), attr.GetPrimitiveType()}
			attrs = append(attrs, a)
		}
		cls = append(cls, class{e.GetName(), attrs})
	}
	scls := superclass{c.GetName(), cls}

	fmt.Println(scls.toString())
}

type superclass struct {
	name string
	class []class
}

func (c *superclass) toString() string {
	str := "/**" + "\n" +
		" * " + c.name + " auto generated code to make your mind swiwel in joy" + "\n" +
		"*/" + "\n"
	for _, cls := range c.class {
		str += cls.toString()
	}
	return str
}

type class struct {
	name string
	attributes []attribute
}

func (c *class) toString() string {
	str := "/**" + "\n" +
		" * " + c.name + "\n" +
		"*/" + "\n" +
		"type " + strings.Title(c.name) + " struct {\n"
	for _, attr := range c.attributes {
		str += attr.toString()
	}
	str += "\n}\n"
	return str
}

type attribute struct {
	name string
	typ string
}

func (c *attribute) toString() string {
	return "\t" + c.name + " " + c.typ + "\n"

}