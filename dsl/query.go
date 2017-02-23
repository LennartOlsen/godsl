package dsl

type Comparable interface {}

type Comparator interface {}

type Query struct {}

func (q Query) Where(comparable Comparable, comparator Comparator, result string) {
	
}