package architecture_class

import (
	"fmt"
)

// Person is how the architecture package stores a person
type Person struct {
	First string
}

// Accessor is how to store or retrieve a person
// When retrieving a person, if they do not exist returns the zero value
type Accessor interface {
	Save(int, Person)
	Retrieve(int) Person
}

type PersonService struct {
	a Accessor
}

func NewPersonService(a Accessor) PersonService {
	return PersonService{a: a}
}

func (ps PersonService) Get(n int) (Person, error) {
	p := ps.a.Retrieve(n)
	if p.First == "" {
		return Person{}, fmt.Errorf("No person found with n of %d", n)
	}
	return p, nil
}

func Put(a Accessor, n int, p Person) {
	a.Save(n, p)
}
func Get(a Accessor, n int) Person {
	return a.Retrieve(n)
}
