package persons

import dsl "github.com/lennartolsen/godsl/query"

type Persons struct {
	dsl.Collection
}

/** EXAMPLES OF OVERRIDING PARENT FUNCTIONS **/
func (c *Persons) Entity(name string) dsl.IEntity {
	ent := c.Collection.Entity(name)
	return ent
}

func (c *Persons) GetType() string {
	return "Persons"
}