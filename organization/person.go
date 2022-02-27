package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// type alias
// a type alias copies the fields and the methods
// type TwitterHandler = string

// type declaration
// a type declaration copies the fields of a type to a new type

// type Handler struct {
// 	handle string
// 	name   string
// }

// func (h Handler) randomFunc() {
// }

type TwitterHandler string

func (twitterHandler TwitterHandler) RedirectURL() string {
	cleanHandler := strings.TrimPrefix(string(twitterHandler), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

type europeanUnionIdentifier struct {
	id      string
	country string
}

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value)
}

func NewEuropeanUnionIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return europeanUnionIdentifier{
			id:      v,
			country: country,
		}
	case int:
		return europeanUnionIdentifier{
			id:      strconv.Itoa(v),
			country: country,
		}
	case europeanUnionIdentifier:
		return v
	case Person:
		euID, _ := v.Citizen.(europeanUnionIdentifier)
		return euID
	default:
		panic("using an invalid type to initialize EU identifier")
	}
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (eui europeanUnionIdentifier) ID() string {
	return eui.id
}

func (ssn socialSecurityNumber) Country() string {
	return "USA"
}

func (eui europeanUnionIdentifier) Country() string {
	return eui.country
}

type Name struct {
	first string
	last  string
}

type Employee struct {
	Name
}

// To make fields public, use Pascal case
// To make fields private, use camel case
type Person struct {
	Name
	twitterHandler TwitterHandler // type alias is a reference to another type
	Citizen                       // embedding interfaces automatically gives access to interface methods
}

// To create an instance of Person
func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.first, p.last)
}

// Person implicitly implements Identifiable interface
func (p *Person) ID() string {
	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
}

// When editing state of a custom type, either use pointer-based receiver or return a new instance of the type
// Otherwise, the state of another copy is actually being changed
func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with @")
	}
	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
