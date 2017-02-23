package classroom

import dsl "github.com/lennartolsen/godsl/query"

type Classroom struct {
	dsl.Collection
}

/** EXAMPLES OF OVERRIDING PARENT FUNCTIONS **/

func (c *Classroom) GetType() string {
	return "Classroom"
}