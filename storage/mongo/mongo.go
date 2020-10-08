package mongo

import (
	architecture_class "github.com/smakaroni/architecture-class"
)

type Mongo map[int]architecture_class.Person

func (m Mongo) Save(i int, p architecture_class.Person) {
	m[i] = p
}

func (m Mongo) Retrieve(i int) architecture_class.Person {
	return m[i]
}
