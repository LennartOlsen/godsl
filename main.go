package main

import (
	dsl "github.com/lennartolsen/godsl/query"
	"fmt"
	"github.com/lennartolsen/godsl/classroom"
)

func main() {
	dsl.RegisterCollectionType("classroom", classroom.Classroom{})
	dsl.RegisterEntityType("teacher", classroom.Teacher{})
	dsl.RegisterEntityType("student", classroom.Student{})
	/**
	 * Consider making generic attribute types, like string, int, float etc in the dsl
	 * More complex types could reside in the persons instance
	 */
	dsl.RegisterAttributeType("name", classroom.String{})
	dsl.RegisterAttributeType("age", classroom.Integer64{})

	classroom := dsl.NewCollection("classroom").
				Entity("teacher").
					Attribute("name").
					Attribute("age").
				Entity("student").
					Attribute("name").
					Attribute("age").
		Flush()

	classroom.
		Append("teacher",
			dsl.V{"name", "lennart"},
			dsl.V{"age", 17},
			dsl.V{"grade", 4}).
		Append("student",
			dsl.V{"name", "lennart"})

	str := classroom.ToString()

	fmt.Println(str)
}

type val struct {
	name string
	value interface{}
}