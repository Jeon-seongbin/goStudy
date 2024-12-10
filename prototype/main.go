package main

import (
	"bytes"
	"encoding/gob"
)

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func main() {

}

func (a *Address) DeepCopy() *Address {
	return &Address{
		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

func (p *Person) DeepCopy() *Person {
	// q := *p
	// q.Address = p.Address.DeepCopy()
	// q.Friends = make([]string, len(p.Friends))
	// copy(q.Friends, p.Friends)

	// return &q
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	if err := encoder.Encode(p); err != nil {
		panic(err)
	}

	decoder := gob.NewDecoder(&buf)
	var result Person
	if err := decoder.Decode(&result); err != nil {
		panic(err)
	}
	return &result
}

func NewPerson(p *Person, name, streetAddress string) *Person {
	q := p.DeepCopy()
	q.Name = name
	q.Address.StreetAddress = streetAddress
	return q
}
