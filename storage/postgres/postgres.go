package postgres

import (
	architecture_class "github.com/smakaroni/architecture-class"
)

type Postgres map[int]architecture_class.Person

func (m Postgres) Save(i int, p architecture_class.Person) {
	m[i] = p
}

func (m Postgres) Retrieve(i int) architecture_class.Person {
	return m[i]
}
