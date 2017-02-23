package persons

import dsl "github.com/lennartolsen/godsl/query"

type Person struct {
	dsl.Ent
}

func (p *Person) Attribute(name string) dsl.IEntity {
	return p.Ent.Attribute(name)
}

func (p *Person) GetType() string {
	return "Person";
}