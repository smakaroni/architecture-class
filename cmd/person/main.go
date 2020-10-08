package main

import (
	"fmt"
	architecture_class "github.com/smakaroni/architecture-class"
	"github.com/smakaroni/architecture-class/storage/mongo"
	"github.com/smakaroni/architecture-class/storage/postgres"
)

func main() {
	dbm := mongo.Mongo{}
	dbp := postgres.Postgres{}

	p1 := architecture_class.Person{
		First: "Emma",
	}
	p2 := architecture_class.Person{First: "Jokke"}
	ps := architecture_class.NewPersonService(dbm)

	architecture_class.Put(dbm, 1, p1)
	architecture_class.Put(dbm, 2, p2)
	fmt.Println(architecture_class.Get(dbm, 1))
	fmt.Println(architecture_class.Get(dbm, 2))
	fmt.Println("---------------------")
	fmt.Println(ps.Get(1))
	fmt.Println(ps.Get(3))
	fmt.Println("---------------------")
	architecture_class.Put(dbp, 1, p1)
	architecture_class.Put(dbp, 2, p2)
	fmt.Println(architecture_class.Get(dbp, 1))
	fmt.Println(architecture_class.Get(dbp, 2))
}
