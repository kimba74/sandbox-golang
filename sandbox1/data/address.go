package data

import "fmt"

type Address struct {
	number              int
	street, city, state string
	zip                 int
}

func (a *Address) GetNumber() int {
	return a.number
}

func (a *Address) GetStreet() string {
	return a.street
}

func (a *Address) GetCity() string {
	return a.city
}

func (a *Address) GetState() string {
	return a.state
}

func (a *Address) GetZip() int {
	return a.zip
}

func (a *Address) String() string {
	if a != nil {
		return fmt.Sprintf("%d %s %s, %s %d", a.number, a.street, a.city, a.state, a.zip)
	}
	return "<nil>"
}

func Create(number int, street, city, state string, zip int) *Address {
	return &Address{number, street, city, state, zip}
}
