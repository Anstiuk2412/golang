package main

import (
	"fmt"
)

type Passport struct {
	Name     string
	Surname  string
	Birthday int
	CodeDoc  Code
}

type Code struct {
	ID string
}

func (doc *Passport) showID() {
	fmt.Println(doc.CodeDoc.ID)
}
func main() {

	doc1 := Passport{
		Name:     "Daniil",
		Surname:  "Anastiuk",
		Birthday: 24122000,
		CodeDoc: Code{
			ID: "7373192394",
		},
	}
	doc2 := Passport{
		Name:     "Jess",
		Surname:  "Rubic",
		Birthday: 12022011,
		CodeDoc: Code{
			ID: "1231321394",
		},
	}

	fmt.Println(doc1)
	fmt.Println(doc2)
	doc1.showID()
}
